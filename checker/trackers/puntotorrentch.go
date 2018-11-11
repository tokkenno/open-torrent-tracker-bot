package trackers

import (
	"github.com/tokkenno/open-torrent-tracker-bot/checker/check"
	"github.com/tokkenno/open-torrent-tracker-bot/models"
	"strings"
)

type PuntotorrentCh struct {
}

func (tracker *PuntotorrentCh) GetId() string {
	return "puntotorrent_ch"
}

func (tracker *PuntotorrentCh) GetName() string {
	return "PuntoTorrent"
}

func (tracker *PuntotorrentCh) GetDescription() string {
	return ""
}

func (tracker *PuntotorrentCh) GetLanguages() []string {
	return []string{"es"}
}

func (tracker *PuntotorrentCh) GetCategories() []string {
	return []string{"movies"}
}

func (tracker *PuntotorrentCh) GetUrl() string {
	return "https://xbt.puntotorrent.ch"
}

func (tracker *PuntotorrentCh) GetRegistryUrl() string {
	return "https://xbt.puntotorrent.ch/index.php?page=signup"
}

func (tracker *PuntotorrentCh) Check() *check.Result {
	html, err := check.GetWebCode(tracker.GetRegistryUrl())

	if err != nil {
		return &check.Result{
			IsOnline: false,
		}
	} else {
		var status models.RegistrationStatus

		if strings.Contains(html, "los registros están cerrados") {
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
