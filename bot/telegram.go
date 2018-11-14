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
	contexts map[int64]string
}

var botInstance *bot

func GetInstance() *bot {
	if botInstance == nil {
		botInstance = &bot{contexts: make(map[int64]string)}
	}
	return botInstance
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

	if len(token) != 45 {
		return errors.New("the telegram bot can't start if the token is undefined or invalid")
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
					case "cancel":
						bot.handleCancel(update.Message)
						break
					default:
						bot.handleUnknown(update.Message)
						break
					}
				} else {
					context, hasContext := bot.contexts[update.Message.Chat.ID]

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

func (bot *bot) SendMessage(id int64, message string, languageCode string) {
	msg := tgbotapi.NewMessage(id, message)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.api.Send(msg)
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

func (bot *bot) handleCancel(message *tgbotapi.Message) {
	bot.contexts[message.Chat.ID] = ""

	msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseActionsCanceled))
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.api.Send(msg)
}

func (bot *bot) handleSubscribe(message *tgbotapi.Message) {
	if len(message.Command()) > 0 {
		bot.contexts[message.Chat.ID] = "subscribe"

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
		bot.contexts[message.Chat.ID] = ""
		msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseIDontUnderstand))
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.api.Send(msg)
	}
}

func (bot *bot) handleSubscribeLanguage(message *tgbotapi.Message) {
	bot.contexts[message.Chat.ID] = "subscribe_language_code"
	msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseInsertLanguage)+strings.Join(checker.GetManager().ListLanguages(), ", "))
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.api.Send(msg)
}

func (bot *bot) handleSubscribeLanguageCode(message *tgbotapi.Message) {
	languagesAvailable := checker.GetManager().ListLanguages()
	languageCode := strings.Split(message.Text, " ")[0]

	if len(languageCode) == 3 && utils.StringSliceContains(languagesAvailable, strings.TrimLeft(languageCode, "-")) {
		languageCode = strings.TrimLeft(languageCode, "-")

		dbRepo := database.GetRepository(models.UserRepositoryName)

		userDocument := models.User{}
		err := dbRepo.Find(bson.M{"telegram_id": message.Chat.ID}).One(&userDocument)

		if err == nil {
			userDocument.Languages = utils.StringSliceFilter(userDocument.Languages, languageCode)
			err = dbRepo.UpdateId(userDocument.Id, userDocument)

			if err != nil {
				logger.Errorf("Error while update the data of the user: %v => %s", userDocument.TelegramId, userDocument.Name)
				bot.handleInternalError(message)
			} else {
				msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseLanguageUnsubscribed))
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				bot.api.Send(msg)
			}
		} else {
			msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseNoSubscribed))
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.api.Send(msg)
		}
	} else if utils.StringSliceContains(languagesAvailable, languageCode) {
		dbRepo := database.GetRepository(models.UserRepositoryName)

		userDocument := models.User{}
		err := dbRepo.Find(bson.M{"telegram_id": message.Chat.ID}).One(&userDocument)

		if err != nil {
			userDocument = models.User{
				TelegramId: message.Chat.ID,
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
			userDocument.Languages = utils.StringUnique(append(userDocument.Languages, languageCode))

			err = dbRepo.UpdateId(userDocument.Id, userDocument)

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
	bot.contexts[message.Chat.ID] = "subscribe_category_code"
	msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseInsertCategory)+strings.Join(checker.GetManager().ListCategories(), ", "))
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.api.Send(msg)
}

func (bot *bot) handleSubscribeCategoryCode(message *tgbotapi.Message) {
	categoriesAvailable := checker.GetManager().ListCategories()
	categoryCode := strings.Split(message.Text, " ")[0]

	if strings.Index(categoryCode, "-") == 0 && utils.StringSliceContains(categoriesAvailable, strings.TrimLeft(categoryCode, "-")) {
		categoryCode = strings.TrimLeft(categoryCode, "-")

		dbRepo := database.GetRepository(models.UserRepositoryName)

		userDocument := models.User{}
		err := dbRepo.Find(bson.M{"telegram_id": message.Chat.ID}).One(&userDocument)

		if err == nil {
			userDocument.Categories = utils.StringSliceFilter(userDocument.Categories, categoryCode)
			err = dbRepo.UpdateId(userDocument.Id, userDocument)

			if err != nil {
				logger.Errorf("Error while update the data of the user: %v => %s", userDocument.TelegramId, userDocument.Name)
				bot.handleInternalError(message)
			} else {
				msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseCategoryUnsubscribed))
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				bot.api.Send(msg)
			}
		} else {
			msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseNoSubscribed))
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.api.Send(msg)
		}
	} else if utils.StringSliceContains(categoriesAvailable, categoryCode) {
		dbRepo := database.GetRepository(models.UserRepositoryName)

		userDocument := models.User{}
		err := dbRepo.Find(bson.M{"telegram_id": message.Chat.ID}).One(&userDocument)

		if err != nil {
			userDocument = models.User{
				TelegramId: message.Chat.ID,
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
			userDocument.Categories = utils.StringUnique(append(userDocument.Categories, categoryCode))

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
	err := dbRepo.Find(bson.M{"telegram_id": message.Chat.ID}).One(&userDocument)

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
