package check

import (
	"errors"
	"github.com/tokkenno/open-torrent-tracker-bot/models"
)

type Tracker interface {
	GetId() string
	GetName() string
	GetDescription() string
	GetLanguages() []string
	GetCategories() []string
	GetUrl() string
	GetRegistryUrl() string
	Check() *Result
}

func CheckerToModel(tracker Tracker) (*models.Tracker, error) {
	if tracker == nil {
		return nil, errors.New("tracker checker uninitialized")
	} else {
		return &models.Tracker{
			FileId:      tracker.GetId(),
			Name:        tracker.GetName(),
			Description: tracker.GetDescription(),
			Languages:   tracker.GetLanguages(),
			Categories:  tracker.GetCategories(),
			OpenStatus:  models.RegistrationClose,
		}, nil
	}
}
