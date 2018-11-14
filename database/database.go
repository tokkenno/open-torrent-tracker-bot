package database

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/google/logger"
	"log"
	"os"
	"strings"
)

const (
	dbName = "torrent-tracker-bot"
)

var instance *mgo.Database

func newDatabase() *mgo.Database {
	host := "localhost"
	port := "27017"
	user := ""
	pass := ""

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == "MONGODB_HOST" {
			host = pair[1]
		} else if pair[0] == "MONGODB_PORT" {
			port = pair[1]
		} else if pair[0] == "MONGODB_USER" {
			user = pair[1]
		} else if pair[0] == "MONGODB_PASS" {
			pass = pair[1]
		}
	}

	if len(user) > 0 && len(pass) > 0 {
		user = fmt.Sprintf("%s:%s@", user, pass)
	} else {
		user = ""
	}

	connectUrl := fmt.Sprintf("mongodb://%s%s:%s", user, host, port)
	logger.Infof("Connecting to database: %s", connectUrl)

	session, err := mgo.Dial(connectUrl)
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