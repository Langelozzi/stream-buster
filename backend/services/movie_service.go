package services

import (
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/models/api"
)

type MovieService struct {
	dao iDao.MovieDatabaseDaoInterface
}

func NewMovieService(dao iDao.MovieDatabaseDaoInterface) *MovieService {
	return &MovieService{dao: dao}
}

func (service *MovieService) GetMovieDetails(id int) (*api.Movie, error) {
	return service.dao.GetMovieDetails(id)
}
