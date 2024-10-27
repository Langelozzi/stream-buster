package daos

import (
	"github.com/STREAM-BUSTER/stream-buster/models/db"
	"github.com/STREAM-BUSTER/stream-buster/utils/database"
)

type CurrentlyWatchingDao struct{}

func NewCurrentlyWatchingDao() *CurrentlyWatchingDao {
	return &CurrentlyWatchingDao{}
}

// CreateCurrentlyWatching creates a new CurrentlyWatching record
func (dao *CurrentlyWatchingDao) CreateCurrentlyWatching(watch *db.CurrentlyWatching) (*db.CurrentlyWatching, error) {
	db := database.GetInstance()

	if err := db.Create(watch).Error; err != nil {
		return nil, err
	}

	return watch, nil
}

// GetCurrentlyWatchingById retrieves a CurrentlyWatching record by its UserID and MediaId
func (dao *CurrentlyWatchingDao) GetCurrentlyWatchingById(userID uint, mediaId string, includeDeleted bool) (*db.CurrentlyWatching, error) {
	databaseInstance := database.GetInstance()

	var currentlyWatching db.CurrentlyWatching
	query := databaseInstance.Model(&db.CurrentlyWatching{}).Where("user_id = ? AND media_id = ?", userID, mediaId)

	if !includeDeleted {
		query.Where("deleted_at IS NULL")
	}

	if err := query.First(&currentlyWatching).Error; err != nil {
		return nil, err
	}

	return &currentlyWatching, nil
}

// UpdateCurrentlyWatching updates the details of an existing CurrentlyWatching record
func (dao *CurrentlyWatchingDao) UpdateCurrentlyWatching(updatedWatch *db.CurrentlyWatching) (*db.CurrentlyWatching, error) {
	databaseInstance := database.GetInstance()

	var existingWatch db.CurrentlyWatching
	if err := databaseInstance.Where("user_id = ? AND media_id = ?", updatedWatch.UserID, updatedWatch.MediaID).First(&existingWatch).Error; err != nil {
		return nil, err
	}

	if err := databaseInstance.Model(&existingWatch).Omit("UserID", "MediaId").Updates(updatedWatch).Error; err != nil {
		return nil, err
	}

	return &existingWatch, nil
}
