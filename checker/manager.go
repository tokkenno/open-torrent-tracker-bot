package checker

import (
	"github.com/globalsign/mgo/bson"
	"github.com/google/logger"
	"github.com/tokkenno/open-torrent-tracker-bot/checker/check"
	"github.com/tokkenno/open-torrent-tracker-bot/checker/trackers"
	"github.com/tokkenno/open-torrent-tracker-bot/database"
	"github.com/tokkenno/open-torrent-tracker-bot/models"
	"github.com/tokkenno/open-torrent-tracker-bot/utils"
	"time"
)

type manager struct {
	trackers         []check.Tracker
	checkInterval chan bool
}

var instance *manager

func GetManager() *manager {
	if instance == nil {
		instance = newManager()
	}
	return instance
}

func newManager() *manager {
	manager := &manager{
		trackers: []check.Tracker{},
	}

	// Spanish
	manager.trackers = append(manager.trackers, []check.Tracker{
		&trackers.DivTeamCom{},
		&trackers.HachedeMe{},
		&trackers.HDBytesLi{},
		&trackers.HDCityLi{},
		&trackers.PuntotorrentCh{},
		&trackers.TbPlusLi{},
		&trackers.TorrentLandLi{},
		&trackers.XBytesV2{},
	}...)

	return manager
}

func (manager *manager) runTrackerCheck(tracker check.Tracker) {
	logger.Infof("Check of the tracker %v started.", tracker.GetName())

	dbRepo := database.GetRepository(models.TrackerRepositoryName)
	resultChannel := make(chan *check.Result)

	go func() {
		resultChannel <- tracker.Check()
	}()

	isNew := false
	trackerDoc := &models.Tracker{}
	err := dbRepo.Find(bson.M{"file_id": tracker.GetId()}).One(&trackerDoc)

	if err != nil {
		trackerDoc, _ = check.CheckerToModel(tracker)
		isNew = true
	}

	trackerDoc.LastCheck = time.Now()
	result := <-resultChannel

	if result == nil {
		logger.Errorf("The check result for the tracker %v is invalid", tracker.GetName())
		return
	}

	if result.IsOnline {
		trackerDoc.LastOnline = time.Now()
		trackerDoc.OpenStatus = result.Status

		if result.Status == models.RegistrationOpen {
			trackerDoc.LastOpen = time.Now()
		}
	}

	if isNew {
		err = dbRepo.Insert(trackerDoc)

		if err != nil {
			logger.Errorf("Error while create the new %v tracker: %v", tracker.GetName(), err)
		}
	} else {
		_, err = dbRepo.UpsertId(trackerDoc.Id, trackerDoc)

		if err != nil {
			logger.Errorf("Error while try to update %v tracker: %v", tracker.GetName(), err)
		}
	}

	logger.Infof("Check of the tracker %v finished.", tracker.GetName())
}

func (manager *manager) RunCheck() {
	logger.Infof("Updating %v trackers...", len(manager.trackers))

	for _, tracker := range manager.trackers {
		manager.runTrackerCheck(tracker)
	}

	logger.Infof("Tracker updating finished.")
}

func (manager *manager) RunIntervalCheck(duration time.Duration) {
	if manager.checkInterval != nil {
		manager.checkInterval <- true
	}

	manager.checkInterval = utils.SetInterval(manager.RunCheck, duration, true)
}

func (manager *manager) ClearIntervalCheck() {
	if manager.checkInterval != nil {
		manager.checkInterval <- true
	}
}

func (manager *manager) ListLanguages() []string {
	var languages []string

	for _, tracker := range manager.trackers {
		languages = append(languages, tracker.GetLanguages()...)
	}

	return utils.StringUnique(languages)
}

func (manager *manager) ListCategories() []string {
	var categories []string

	for _, tracker := range manager.trackers {
		categories = append(categories, tracker.GetCategories()...)
	}

	return utils.StringUnique(categories)
}