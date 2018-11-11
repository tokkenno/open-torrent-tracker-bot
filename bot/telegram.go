package bot

import (
	"errors"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tokkenno/open-torrent-tracker-bot/checker"
	"github.com/tokkenno/open-torrent-tracker-bot/language"
	"log"
	"os"
	"strings"
)

type bot struct {
	api *tgbotapi.BotAPI
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
	msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseInsertLanguage) + strings.Join(checker.GetManager().ListLanguages(), ", "))
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.api.Send(msg)
}

func (bot *bot) handleSubscribeLanguageCode(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseLanguageSubscribed))
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.api.Send(msg)
}

func (bot *bot) handleSubscribeCategory(message *tgbotapi.Message) {
	bot.contexts[message.From.ID] = "subscribe_category_code"
	msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseInsertCategory) + strings.Join(checker.GetManager().ListCategories(), ", "))
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.api.Send(msg)
}

func (bot *bot) handleSubscribeCategoryCode(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseCategorySubscribed))
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.api.Send(msg)
}

func (bot *bot) handleSubscriptions(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, language.Localize(message.From.LanguageCode, language.PhraseInsertLanguage))
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.api.Send(msg)
}