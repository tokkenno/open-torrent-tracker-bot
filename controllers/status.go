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

type status struct {
	Tracker statusTrackers `json:"tracker"`
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	dbInstance := database.GetRepository(models.TrackerRepositoryName)

	var trackers []models.Tracker
	err := dbInstance.Find(nil).All(&trackers)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
		return
	}

	status := status{
		Tracker: statusTrackers{
			Count: len(trackers),
		},
	}

	statusJson, _ := json.Marshal(status)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(statusJson))
}
