package models

import (
	"github.com/globalsign/mgo/bson"
)

type User struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	TelegramId uint64        `bson:"telegram_id"`
	Languages  []string      `bson:"languages"`
	Categories []string      `bson:"categories"`
}
