package services

import (
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/models/db"
)

type MediaService struct {
	dao iDao.MediaDaoInterface
}

func NewMediaService(dao iDao.MediaDaoInterface) *MediaService {
	return &MediaService{dao: dao}
}
func (service MediaService) GetMediaById(id int64) (*db.Media, error) {
	return service.dao.GetMediaById(id)
}
func (service MediaService) CreateMedia(media *db.Media) error {
	return service.dao.CreateMedia(media)
}
