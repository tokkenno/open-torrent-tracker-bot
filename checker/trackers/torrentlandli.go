package trackers

import (
	"github.com/tokkenno/open-torrent-tracker-bot/checker/check"
	"github.com/tokkenno/open-torrent-tracker-bot/models"
	"strings"
)

type TorrentLandLi struct {
}

func (tracker *TorrentLandLi) GetId() string {
	return "torrentland_li"
}

func (tracker *TorrentLandLi) GetName() string {
	return "Torrent Land"
}

func (tracker *TorrentLandLi) GetDescription() string {
	return ""
}

func (tracker *TorrentLandLi) GetLanguages() []string {
	return []string{"es"}
}

func (tracker *TorrentLandLi) GetCategories() []string {
	return []string{"movies"}
}

func (tracker *TorrentLandLi) GetUrl() string {
	return "http://torrentland.li"
}

func (tracker *TorrentLandLi) GetRegistryUrl() string {
	return "http://torrentland.li/sbg_login_classic.php"
}

func (tracker *TorrentLandLi) Check() *check.Result {
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
