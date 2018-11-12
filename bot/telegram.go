package bot

import (
	"errors"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/logger"
	"github.com/tokkenno/open-torrent-tracker-bot/checker"
	"github.com/tokkenno/open-torrent-tracker-bot/database"
	"github.com/tokkenno/open-torrent-tracker-bot/language"
	"github.com/tokkenno/open-torrent-tracker-bot/models"
	"github.com/tokkenno/open-torrent-tracker-bot/utils"
	"log"
	"os"
	"strings"
)

type bot struct {
	api      *tgbotapi.BotAPI
	contexts map[int]string
}

func NewTelegramBot() *bot {
	return &bot{contexts: make(map[int]string)}
}

func (bot *bot) ListenAsync() error {
	debug := false
	var token string

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == "TELEGRAM_TOKEN" {
			token = pair[1]
		} else if pair[0] == "BOT_DEBUG" {
			debug = true
		}
	}

	if len(token) == 0 {
		return errors.New("the telegram bot can't start if the token is undefined")
	}

	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.api = botApi
	bot.api.Debug = debug

	log.Printf("Authorized on account %s", bot.api.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	go func() {
		updates, _ := bot.api.GetUpdatesChan(u)

		for update := range updates {
			if update.Message != nil {
				if len(update.Message.Command()) > 0 {
					switch update.Message.Command() {
					case "start":
						bot.handleStart(update.Message)
						break
					case "info":
						bot.handleInfo(update.Message)
						break
					case "subscribe":
						bot.handleSubscribe(update.Message)
						break
					case "subscriptions":
						bot.handleSubscriptions(update.Message)
						break
					default:
						bot.handleUnknown(update.Message)
						break
					}
				} else {
					context, hasContext := bot.contexts[update.Message.From.ID]

					if hasContext && len(context) > 0 {
						switch context {
						case "subscribe":
							bot.handleSubscribe(update.Message)
						case "subscribe_language_code":
							bot.handleSubscribeLanguageCode(update.Message)
						case "subscribe_category_code":
							bot.handleSubscribeCategoryCode(update.Message)
						}
					} else {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, language.Localize(update.Message.From.LanguageCode, language.PhraseIDontUnderstand))
						msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
						bot.api.Send(msg)
					}
				}
			}
		}
	}()

	return nil
}

func (bot *bot) handleStart(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.MessageCommandStart))
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.api.Send(msg)
}

func (bot *bot) handleInfo(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.MessageCommandInfo))
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.api.Send(msg)
}

func (bot *bot) handleUnknown(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.MessageCommandUnknown))
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.api.Send(msg)
}

func (bot *bot) handleSubscribe(message *tgbotapi.Message) {
	if len(message.Command()) > 0 {
		bot.contexts[message.From.ID] = "subscribe"

		subscribeKeyboard := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(language.Localize(message.From.LanguageCode, language.WordCategory)),
				tgbotapi.NewKeyboardButton(language.Localize(message.From.LanguageCode, language.WordLanguage)),
			),
		)

		msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseChooseSubscription))
		msg.ReplyMarkup = subscribeKeyboard
		bot.api.Send(msg)
	} else if message.Text == language.Localize(message.From.LanguageCode, language.WordCategory) {
		bot.handleSubscribeCategory(message)
	} else if message.Text == language.Localize(message.From.LanguageCode, language.WordLanguage) {
		bot.handleSubscribeLanguage(message)
	} else {
		bot.contexts[message.From.ID] = ""
		msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseIDontUnderstand))
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.api.Send(msg)
	}
}

func (bot *bot) handleSubscribeLanguage(message *tgbotapi.Message) {
	bot.contexts[message.From.ID] = "subscribe_language_code"
	msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseInsertLanguage)+strings.Join(checker.GetManager().ListLanguages(), ", "))
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.api.Send(msg)
}

func (bot *bot) handleSubscribeLanguageCode(message *tgbotapi.Message) {
	lenguagesAvailable := checker.GetManager().ListLanguages()
	languageCode := strings.Split(message.Text, " ")[0]

	if utils.StringSliceContains(lenguagesAvailable, languageCode) {
		dbRepo := database.GetRepository(models.UserRepositoryName)

		userDocument := models.User{}
		err := dbRepo.Find(bson.M{"telegram_id": message.From.ID}).One(&userDocument)

		if err != nil {
			userDocument = models.User{
				TelegramId: message.From.ID,
				Name:       message.From.UserName,
				Categories: []string{},
				Languages:  []string{languageCode},
			}

			err = dbRepo.Insert(userDocument)

			if err != nil {
				logger.Errorf("Error while create the new user: %v => %s", userDocument.TelegramId, userDocument.Name)
				bot.handleInternalError(message)
				return
			}
		} else {
			userDocument.Languages = append(userDocument.Languages, languageCode)

			_, err = dbRepo.UpsertId(userDocument.Id, userDocument)

			if err != nil {
				logger.Errorf("Error while update the data of the user: %v => %s", userDocument.TelegramId, userDocument.Name)
				bot.handleInternalError(message)
				return
			}
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseLanguageSubscribed))
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.api.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseLanguageUnrecognized))
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.api.Send(msg)
	}
}

func (bot *bot) handleSubscribeCategory(message *tgbotapi.Message) {
	bot.contexts[message.From.ID] = "subscribe_category_code"
	msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseInsertCategory)+strings.Join(checker.GetManager().ListCategories(), ", "))
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.api.Send(msg)
}

func (bot *bot) handleSubscribeCategoryCode(message *tgbotapi.Message) {
	categoriesAvailable := checker.GetManager().ListCategories()
	categoryCode := strings.Split(message.Text, " ")[0]

	if utils.StringSliceContains(categoriesAvailable, categoryCode) {
		dbRepo := database.GetRepository(models.UserRepositoryName)

		userDocument := models.User{}
		err := dbRepo.Find(bson.M{"telegram_id": message.From.ID}).One(&userDocument)

		if err != nil {
			userDocument = models.User{
				TelegramId: message.From.ID,
				Name:       message.From.UserName,
				Categories: []string{categoryCode},
				Languages:  []string{},
			}

			err = dbRepo.Insert(userDocument)

			if err != nil {
				logger.Errorf("Error while create the new user: %v => %s", userDocument.TelegramId, userDocument.Name)
				bot.handleInternalError(message)
				return
			}
		} else {
			userDocument.Categories = append(userDocument.Categories, categoryCode)

			_, err = dbRepo.UpsertId(userDocument.Id, userDocument)

			if err != nil {
				logger.Errorf("Error while update the data of the user: %v => %s", userDocument.TelegramId, userDocument.Name)
				bot.handleInternalError(message)
				return
			}
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseCategorySubscribed))
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.api.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseCategoryUnrecognized))
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.api.Send(msg)
	}
}

func (bot *bot) handleSubscriptions(message *tgbotapi.Message) {
	dbRepo := database.GetRepository(models.UserRepositoryName)

	userDocument := models.User{}
	err := dbRepo.Find(bson.M{"telegram_id": message.From.ID}).One(&userDocument)

	if err == nil && userDocument.Categories != nil && len(userDocument.Categories) > 0 || userDocument.Languages != nil && len(userDocument.Languages) > 0 {
		if err == nil && userDocument.Categories != nil && len(userDocument.Categories) > 0 {
			msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(language.Localize(message.From.LanguageCode, language.PhraseUserSubscriptionsCategories), strings.Join(userDocument.Categories, ", ")))
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.api.Send(msg)
		}

		if userDocument.Languages != nil && len(userDocument.Languages) > 0 {
			msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(language.Localize(message.From.LanguageCode, language.PhraseUserSubscriptionsLanguages), strings.Join(userDocument.Languages, ", ")))
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.api.Send(msg)
		}
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseNoSubscriptions))
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.api.Send(msg)
	}
}

func (bot *bot) handleInternalError(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseInternalError))
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.api.Send(msg)
}
