package trackers

import (
	"github.com/tokkenno/open-torrent-tracker-bot/checker/check"
	"github.com/tokkenno/open-torrent-tracker-bot/models"
	"strings"
)

type XBytesV2 struct {
}

func (tracker *XBytesV2) GetId() string {
	return "xbytesv2_li"
}

func (tracker *XBytesV2) GetName() string {
	return "XBytes (v2)"
}

func (tracker *XBytesV2) GetDescription() string {
	return ""
}

func (tracker *XBytesV2) GetLanguages() []string {
	return []string{"es"}
}

func (tracker *XBytesV2) GetCategories() []string {
	return []string{"movies"}
}

func (tracker *XBytesV2) GetUrl() string {
	return "http://xbytesv2.li"
}

func (tracker *XBytesV2) GetRegistryUrl() string {
	return "http://xbytesv2.li/sbg_login_classic.php"
}

func (tracker *XBytesV2) Check() *check.Result {
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
