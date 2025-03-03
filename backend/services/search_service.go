package services

import (
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/models/api"
)

type SearchService struct {
	dao iDao.MovieDatabaseDaoInterface
}

func NewSearchService(dao iDao.MovieDatabaseDaoInterface) *SearchService {
	return &SearchService{dao: dao}
}

// SearchMultiMedia searches for multimedia content based on a query
func (service *SearchService) SearchMultiMedia(query string, page int) (*api.SearchPage, error) {
	searchResults, err := service.dao.SearchMultiMedia(query, page)
	if err != nil {
		return nil, err
	}

	return searchResults, nil
}

func (service *SearchService) SearchMovies(query string, page int) (*api.SearchPage, error) {
	searchResults, err := service.dao.SearchMovies(query, page)
	if err != nil {
		return nil, err
	}

	return searchResults, nil
}

func (service *SearchService) SearchTV(query string, page int) (*api.SearchPage, error) {
	searchResults, err := service.dao.SearchTV(query, page)
	if err != nil {
		return nil, err
	}

	return searchResults, nil
}

func (service *SearchService) SearchTrendingMovies(timeWindow string, page int) (*api.SearchPage, error) {
	searchResults, err := service.dao.GetTrendingMovies(timeWindow, page)
	if err != nil {
		return nil, err
	}

	return searchResults, nil
}

func (service *SearchService) SearchTrendingTV(timeWindow string, page int) (*api.SearchPage, error) {
	searchResults, err := service.dao.GetTrendingTV(timeWindow, page)
	if err != nil {
		return nil, err
	}

	return searchResults, nil
}
