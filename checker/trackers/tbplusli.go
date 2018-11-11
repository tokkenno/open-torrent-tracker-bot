package trackers

import (
	"github.com/tokkenno/open-torrent-tracker-bot/checker/check"
	"github.com/tokkenno/open-torrent-tracker-bot/models"
	"strings"
)

type TbPlusLi struct {
}

func (tracker *TbPlusLi) GetId() string {
	return "tbplus_li"
}

func (tracker *TbPlusLi) GetName() string {
	return "TBPlus"
}

func (tracker *TbPlusLi) GetDescription() string {
	return ""
}

func (tracker *TbPlusLi) GetLanguages() []string {
	return []string{"es"}
}

func (tracker *TbPlusLi) GetCategories() []string {
	return []string{"movies"}
}

func (tracker *TbPlusLi) GetUrl() string {
	return "https://tbplus.li"
}

func (tracker *TbPlusLi) GetRegistryUrl() string {
	return "https://tbplus.li/index.php?page=signup"
}

func (tracker *TbPlusLi) Check() *check.Result {
	html, err := check.GetWebCode(tracker.GetRegistryUrl())

	if err != nil {
		return &check.Result{
			IsOnline: false,
		}
	} else {
		var status models.RegistrationStatus

		if strings.Contains(html, "los registros estan cerrados") {
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
