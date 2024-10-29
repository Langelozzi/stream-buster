package services

import (
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/models/api"
)

type TVService struct {
	dao iDao.MovieDatabaseDaoInterface
}

func NewTVService(dao iDao.MovieDatabaseDaoInterface) *TVService {
	return &TVService{dao: dao}
}

func (service *TVService) GetTVDetails(id int) (*api.TV, error) {
	return service.dao.GetTVDetails(id)
}

func (service *TVService) GetEpisodesInSeason(seriesId int, seasonNum int) ([]*api.Episode, error) {
	return service.dao.GetEpisodesInSeason(seriesId, seasonNum)
}
