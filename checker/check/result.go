package check

import "github.com/tokkenno/open-torrent-tracker-bot/models"

type Result struct {
	IsOnline bool
	Status   models.RegistrationStatus
}
