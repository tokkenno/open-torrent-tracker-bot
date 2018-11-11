package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type RegistrationStatus int

const (
	RegistrationClose = 0
	RegistrationOpen  = 1
	RegistrationMaybe = 2
)

const (
	TrackerRepositoryName = "tracker"
)

type Tracker struct {
	Id          bson.ObjectId      `bson:"_id,omitempty"`
	FileId      string             `bson:"file_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Languages   []string           `bson:"languages"`
	Categories  []string           `bson:"categories"`
	OpenStatus  RegistrationStatus `bson:"open_status"`
	LastCheck   time.Time          `bson:"last_check"`
	LastOnline  time.Time          `bson:"last_online"`
	LastOpen    time.Time          `bson:"last_open"`
}
