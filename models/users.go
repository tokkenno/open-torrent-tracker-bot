package models

import (
	"github.com/globalsign/mgo/bson"
)

const (
	UserRepositoryName = "user"
)

type User struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	TelegramId   int64         `bson:"telegram_id"`
	Name         string        `bson:"username"`
	UserLanguage string        `bson:"user_language"`
	Languages    []string      `bson:"languages"`
	Categories   []string      `bson:"categories"`
}
