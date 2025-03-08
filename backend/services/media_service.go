package services

import (
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/models/db"
	iServices "github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"strconv"
	"time"
)

type MediaService struct {
	dao        iDao.MediaDaoInterface
	cdnService iServices.CDNServiceInterface
}

func NewMediaService(dao iDao.MediaDaoInterface, cdnService iServices.CDNServiceInterface) *MediaService {
	return &MediaService{
		dao:        dao,
		cdnService: cdnService,
	}
}
func (service MediaService) GetMediaById(id int64) (*db.Media, error) {
	return service.dao.GetMediaById(id)
}
func (service MediaService) GetMediaByTMDBId(id int64) (*db.Media, error) {
	return service.dao.GetMediaByTMDBId(id)
}
func (service MediaService) CreateMedia(media *db.Media) (*db.Media, error) {
	return service.dao.CreateMedia(media)
}

func (service MediaService) GetMediaAvailability(media *db.Media) (int, error) {
	existingMedia, err := service.GetMediaByTMDBId(int64(media.TMDBID))
	if existingMedia != nil {
		media = existingMedia
	}

	if err != nil || shouldRefreshAvailability(media) { // Media doesn't exist yet or availability needs refresh

		// Get the availability
		exists := service.cdnService.CheckContentExists(strconv.Itoa(media.TMDBID), media.MediaTypeId == 1)

		// Create/update media object with availability
		media.Availability = exists
		media, err = service.CreateMedia(media)

		return exists, nil
	}

	// If availability is already stored in database and has been updated within the past 2 weeks
	return media.Availability, nil
}

func shouldRefreshAvailability(media *db.Media) bool {
	twoWeeksAgo := time.Now().AddDate(0, 0, -14) // 14 days ago
	return media.Availability == -1 || (media.UpdatedAt != nil && media.UpdatedAt.Before(twoWeeksAgo))
}
