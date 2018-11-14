package trackers

import (
	"github.com/tokkenno/open-torrent-tracker-bot/checker/check"
	"github.com/tokkenno/open-torrent-tracker-bot/models"
	"strings"
)

type DivTeamCom struct {
}

func (tracker *DivTeamCom) GetId() string {
	return "divteam_com"
}


func (tracker *DivTeamCom) GetName() string {
	return "DivTeam"
}

func (tracker *DivTeamCom) GetDescription() string {
	return ""
}

func (tracker *DivTeamCom) GetLanguages() []string {
	return []string{"es"}
}

func (tracker *DivTeamCom) GetCategories() []string {
	return []string{"movies"}
}

func (tracker *DivTeamCom) GetUrl() string {
	return "https://divteam.com"
}

func (tracker *DivTeamCom) GetRegistryUrl() string {
	return "https://divteam.com/app/"
}

func (tracker *DivTeamCom) Check() *check.Result {
	html, err := check.GetWebCode(tracker.GetRegistryUrl())

	if err != nil {
		return &check.Result{
			IsOnline: false,
		}
	} else {
		var status models.RegistrationStatus

		if strings.Contains(html, "REGISTRO TEMPORALMENTE DESHABILITADO") {
			status = models.RegistrationClose
		} else {
			status = models.RegistrationMaybe
		}
status = models.RegistrationMaybe
		return &check.Result{
			IsOnline: true,
			Status:   status,
		}
	}
}
