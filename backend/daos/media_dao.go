package daos

import (
	"github.com/STREAM-BUSTER/stream-buster/models/db"
	"github.com/STREAM-BUSTER/stream-buster/utils/database"
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

	if err := databaseInstance.Create(media).Error; err != nil {
		return err

	}
	return nil
}
