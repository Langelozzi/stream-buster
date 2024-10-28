package daos

import (
	"github.com/STREAM-BUSTER/stream-buster/models/db"
	"github.com/STREAM-BUSTER/stream-buster/utils/database"
)

type MediaDao struct{}

func NewMediaDao() *MediaDao {
	return &MediaDao{}
}

func (MediaDao) GetMediaById(mediaId uint, includeDeleted bool) (*db.Media, error) {
	databaseInstance := database.GetInstance()
	var media db.Media
	query := databaseInstance.Model(&db.Media{}).Where("id = ?", mediaId)

	if !includeDeleted {
		query.Where("deleted_at IS NULL")
	}

	if err := query.First(&media).Error; err != nil {
		return nil, err
	}

	return &media, nil
}
