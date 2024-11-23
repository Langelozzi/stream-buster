package daos

import (
	"github.com/STREAM-BUSTER/stream-buster/models/db"
	"github.com/STREAM-BUSTER/stream-buster/utils/database"
	"gorm.io/gorm/clause"
)

type CurrentlyWatchingDao struct{}

func NewCurrentlyWatchingDao() *CurrentlyWatchingDao {
	return &CurrentlyWatchingDao{}
}

// CreateCurrentlyWatching creates a new CurrentlyWatching record
func (dao *CurrentlyWatchingDao) CreateCurrentlyWatching(watch *db.CurrentlyWatching) (*db.CurrentlyWatching, error) {
	db := database.GetInstance()

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "media_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"episode_number", "season_number", "updated_at"}),
	}).Create(watch)

	return watch, nil
}

// GetCurrentlyWatchingById retrieves a CurrentlyWatching record by its UserID and MediaId
func (dao *CurrentlyWatchingDao) GetCurrentlyWatchingById(userID uint, mediaId uint, includeDeleted bool) (*db.CurrentlyWatching, error) {
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

func (dao *CurrentlyWatchingDao) GetCurrentlyWatchingByUserId(userID uint, includeDeleted bool) ([]*db.CurrentlyWatching, error) {
	databaseInstance := database.GetInstance()

	var currentlyWatchingRecords []*db.CurrentlyWatching
	query := databaseInstance.Model(&db.CurrentlyWatching{}).Where("user_id = ?", userID)

	if !includeDeleted {
		query = query.Where("deleted_at IS NULL") // Assuming soft delete
	}

	if err := query.Find(&currentlyWatchingRecords).Error; err != nil {
		return nil, err
	}

	return currentlyWatchingRecords, nil
}

func (dao *CurrentlyWatchingDao) GetWatchlist(userId uint) ([]db.CurrentlyWatching, error) {
	databaseInstance := database.GetInstance()

	var watchingList []db.CurrentlyWatching

	// Query CurrentlyWatching records with related Media records for the specified UserID
	err := databaseInstance.Preload("Media.MediaType").Preload("Media").Where("user_id = ?", userId).Find(&watchingList).Error
	if err != nil {
		// Handle error
		return nil, err
	}

	// watchingList now contains all CurrentlyWatching records with the related Media populated
	return watchingList, nil
}

// UpdateCurrentlyWatching updates the details of an existing CurrentlyWatching record
func (dao *CurrentlyWatchingDao) UpdateCurrentlyWatching(updatedWatch *db.CurrentlyWatching) (*db.CurrentlyWatching, error) {
	databaseInstance := database.GetInstance()

	var existingWatch db.CurrentlyWatching
	if err := databaseInstance.Where("user_id = ? AND media_id = ?", updatedWatch.UserID, updatedWatch.MediaId).First(&existingWatch).Error; err != nil {
		return nil, err
	}

	if err := databaseInstance.Model(&existingWatch).Omit("UserID", "MediaId").Updates(updatedWatch).Error; err != nil {
		return nil, err
	}

	return &existingWatch, nil
}

func (dao *CurrentlyWatchingDao) DeleteCurrentlyWatching(userId, mediaId uint) error {
	databaseInstance := database.GetInstance()
	watch := db.CurrentlyWatching{UserID: userId, MediaId: mediaId}

	result := databaseInstance.Delete(&watch)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
