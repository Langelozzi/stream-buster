package daos

import (
	"github.com/STREAM-BUSTER/stream-buster/models/db"
	"github.com/STREAM-BUSTER/stream-buster/utils/database"
	"gorm.io/gorm/clause"
)

type MediaDao struct{}

func NewMediaDao() *MediaDao {
	return &MediaDao{}
}
func (dao MediaDao) GetMediaById(id int64) (*db.Media, error) {
	databaseInstance := database.GetInstance()
	var media db.Media

	query := databaseInstance.Model(&media)

	// if full {
	// 	query.Preload("Usage").Preload("UserRoles.Role").Preload("Configs.Config")
	// }
	// if !includeDeleted {
	// 	query.Where("deleted_at IS NULL")
	// }
	if err := query.First(&media, id).Error; err != nil {
		return nil, err

	}

	return &media, nil
}

func (dao MediaDao) CreateMedia(media *db.Media) error {
	databaseInstance := database.GetInstance()

	if err := databaseInstance.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"}, // Define the unique constraint causing the conflict
		},
		DoUpdates: clause.AssignmentColumns([]string{"overview", "poster_image"}), // Specify columns to update
	}).Create(media).Error; err != nil {
		return err
	}
	return nil
}
