package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/tokkenno/open-torrent-tracker-bot/database"
	"github.com/tokkenno/open-torrent-tracker-bot/models"
	"net/http"
)

type statusTrackers struct {
	Count int `json:"count"`
}

type statusUsers struct {
	Count int `json:"count"`
}

type status struct {
	Trackers statusTrackers `json:"trackers"`
	Users statusUsers `json:"users"`
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	dbTracker := database.GetRepository(models.TrackerRepositoryName)
	dbUser := database.GetRepository(models.UserRepositoryName)

	var trackers []models.Tracker
	err := dbTracker.Find(nil).All(&trackers)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}

	var users []models.User
	err = dbUser.Find(nil).All(&users)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}

	status := status{
		Trackers: statusTrackers{
			Count: len(trackers),
		},
		Users: statusUsers{
			Count: len(users),
		},
	}

	statusJson, _ := json.Marshal(status)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(statusJson))
}
