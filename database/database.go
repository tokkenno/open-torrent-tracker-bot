package database

import (
	"github.com/globalsign/mgo"
	"log"
)

const (
	dbName = "torrent-tracker-bot"
)

var instance *mgo.Database

func newDatabase() *mgo.Database {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session.DB(dbName)
}

func GetDatabase() *mgo.Database {
	if instance == nil {
		instance = newDatabase()
	}
	return instance
}

func GetRepository(name string) *mgo.Collection {
	return GetDatabase().C(name)
}