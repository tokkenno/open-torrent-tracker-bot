package trackers

import (
	"github.com/tokkenno/open-torrent-tracker-bot/checker/check"
	"github.com/tokkenno/open-torrent-tracker-bot/models"
	"strings"
)

type HDBytesLi struct {
}

func (tracker *HDBytesLi) GetId() string {
	return "hdbytes_li"
}

func (tracker *HDBytesLi) GetName() string {
	return "HDBytes"
}

func (tracker *HDBytesLi) GetDescription() string {
	return ""
}

func (tracker *HDBytesLi) GetLanguages() []string {
	return []string{"es"}
}

func (tracker *HDBytesLi) GetCategories() []string {
	return []string{"movies"}
}

func (tracker *HDBytesLi) GetUrl() string {
	return "http://www.hdbytes.li"
}

func (tracker *HDBytesLi) GetRegistryUrl() string {
	return "http://www.hdbytes.li/index.php?page=signup"
}

func (tracker *HDBytesLi) Check() *check.Result {
	html, err := check.GetWebCode(tracker.GetRegistryUrl())

	if err != nil {
		return &check.Result{
			IsOnline: false,
		}
	} else {
		var status models.RegistrationStatus

		if strings.Contains(html, "las inscripciones están cerradas") {
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
