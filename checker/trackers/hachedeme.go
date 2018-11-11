package trackers

import (
	"github.com/tokkenno/open-torrent-tracker-bot/checker/check"
	"github.com/tokkenno/open-torrent-tracker-bot/models"
	"strings"
)

type HachedeMe struct {
}

func (tracker *HachedeMe) GetId() string {
	return "hachede_me"
}

func (tracker *HachedeMe) GetName() string {
	return "+HacheDe"
}

func (tracker *HachedeMe) GetDescription() string {
	return ""
}

func (tracker *HachedeMe) GetLanguages() []string {
	return []string{"es"}
}

func (tracker *HachedeMe) GetCategories() []string {
	return []string{"movies"}
}

func (tracker *HachedeMe) GetUrl() string {
	return "https://hachede.me"
}

func (tracker *HachedeMe) GetRegistryUrl() string {
	return "https://hachede.me/?p=signup&pid=16"
}

func (tracker *HachedeMe) Check() *check.Result {
	html, err := check.GetWebCode(tracker.GetRegistryUrl())

	if err != nil {
		return &check.Result{
			IsOnline: false,
		}
	} else {
		var status models.RegistrationStatus

		if strings.Contains(html, "Lo sentimos, pero en estos momentos los registros están cerrados") {
			status = models.RegistrationClose
		} else {
			status = models.RegistrationMaybe
		}

		return &check.Result{
			IsOnline: true,
			Status:   status,
		}
	}
}
