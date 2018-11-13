package models

import (
	"github.com/globalsign/mgo/bson"
)

const (
	UserRepositoryName = "user"
)

type User struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	TelegramId int           `bson:"telegram_id"`
	Name       string        `bson:"username"`
	Languages  []string      `bson:"languages"`
	Categories []string      `bson:"categories"`
}
