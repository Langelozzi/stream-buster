package services

import (
	"github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/models/db"
)

type CurrentlyWatchingService struct {
	dao interfaces.CurrentlyWatchingDaoInterface
}

// Constructor for CurrentlyWatchingService
func NewCurrentlyWatchingService(dao interfaces.CurrentlyWatchingDaoInterface) *CurrentlyWatchingService {
	return &CurrentlyWatchingService{dao: dao}
}

// Method to create a new CurrentlyWatching entry
func (service *CurrentlyWatchingService) CreateCurrentlyWatching(watch *db.CurrentlyWatching) (*db.CurrentlyWatching, error) {
	return service.dao.CreateCurrentlyWatching(watch)
}

// Method to retrieve a CurrentlyWatching entry by userID and mediaId
func (service *CurrentlyWatchingService) GetCurrentlyWatchingById(userID uint, mediaId uint, includeDeleted bool) (*db.CurrentlyWatching, error) {
	return service.dao.GetCurrentlyWatchingById(userID, mediaId, includeDeleted)
}

// Method to retrieve a CurrentlyWatching entry by userID and mediaId
func (service *CurrentlyWatchingService) GetCurrentlyWatchingByUserId(userID uint, includeDeleted bool) ([]*db.CurrentlyWatching, error) {
	return service.dao.GetCurrentlyWatchingByUserId(userID, includeDeleted)
}

// Method to update an existing CurrentlyWatching entry
func (service *CurrentlyWatchingService) UpdateCurrentlyWatching(updatedWatch *db.CurrentlyWatching) (*db.CurrentlyWatching, error) {
	return service.dao.UpdateCurrentlyWatching(updatedWatch)
}
