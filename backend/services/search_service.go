package services

import (
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
)

type SearchService struct {
	dao iDao.MovieDatabaseDaoInterface
}

func NewSearchService(dao iDao.MovieDatabaseDaoInterface) *SearchService {
	return &SearchService{dao: dao}
}

func (service *SearchService) SearchMultiMedia(query string) ([]interface{}, error) {
	return service.dao.SearchMultiMedia(query)
}
