package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models/db"

type CurrentlyWatchingDaoInterface interface {
	CreateCurrentlyWatching(watch *db.CurrentlyWatching) (*db.CurrentlyWatching, error)
	GetCurrentlyWatchingById(userID uint, mediaId string, includeDeleted bool) (*db.CurrentlyWatching, error)
	UpdateCurrentlyWatching(updatedWatch *db.CurrentlyWatching) (*db.CurrentlyWatching, error)
}
