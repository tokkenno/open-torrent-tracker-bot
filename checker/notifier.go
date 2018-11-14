package checker

type Notifier interface {
	SendMessage(id int64, message string, languageCode string)
}
