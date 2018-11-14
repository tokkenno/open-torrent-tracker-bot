package main

import (
	"bytes"
	"github.com/google/logger"
	"github.com/tokkenno/open-torrent-tracker-bot/bot"
	"github.com/tokkenno/open-torrent-tracker-bot/checker"
	"github.com/tokkenno/open-torrent-tracker-bot/controllers"
	"net/http"
	"time"
)

func main() {
	var buf1 bytes.Buffer
	logger.Init("torrent-tracker-bot", true, false, &buf1)

	botInstance := bot.GetInstance()
	err := botInstance.ListenAsync()

	if err == nil {
		manager := checker.GetManager()
		manager.SetNotifier(botInstance)
		manager.RunIntervalCheck(time.Minute * 15)

		http.HandleFunc("/status", controllers.GetStatus)
		http.ListenAndServe(":80", nil)
	} else {
		logger.Fatal(err)
	}
}
