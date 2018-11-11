package trackers

import (
	"github.com/tokkenno/open-torrent-tracker-bot/checker/check"
	"github.com/tokkenno/open-torrent-tracker-bot/models"
	"strings"
)

type HDCityLi struct {
}

func (tracker *HDCityLi) GetId() string {
	return "hdcity_li"
}

func (tracker *HDCityLi) GetName() string {
	return "HDCity"
}

func (tracker *HDCityLi) GetDescription() string {
	return ""
}

func (tracker *HDCityLi) GetLanguages() []string {
	return []string{"es"}
}

func (tracker *HDCityLi) GetCategories() []string {
	return []string{"movies"}
}

func (tracker *HDCityLi) GetUrl() string {
	return "https://hdcity.li"
}

func (tracker *HDCityLi) GetRegistryUrl() string {
	return "https://hdcity.li/index.php?page=account"
}

func (tracker *HDCityLi) Check() *check.Result {
	html, err := check.GetWebCode(tracker.GetRegistryUrl())

	if err != nil {
		return &check.Result{
			IsOnline: false,
		}
	} else {
		var status models.RegistrationStatus

		if strings.Contains(html, "registrations are closed") {
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
