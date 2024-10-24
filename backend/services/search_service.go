package services

import (
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/models/api"
	"github.com/STREAM-BUSTER/stream-buster/utils"
)

type SearchService struct {
	dao iDao.MovieDatabaseDaoInterface
}

func NewSearchService(dao iDao.MovieDatabaseDaoInterface) *SearchService {
	return &SearchService{dao: dao}
}

func (service *SearchService) SearchMultiMedia(query string) ([]interface{}, error) {
	searchResults, err := service.dao.SearchMultiMedia(query)
	if err != nil {
		return nil, err
	}

	// Filter results to only return the media that has a valid vidsrc url
	validMedia := utils.Filter(searchResults, doesMediaContentExist)

	return validMedia, nil
}

func doesMediaContentExist(item interface{}) bool {
	switch v := item.(type) {
	case *api.TV:
		return utils.DoesContentExistForTMDBId(v.Media.TMDBID)
	case *api.Movie:
		return utils.DoesContentExistForTMDBId(v.Media.TMDBID)
	default:
		return false // or handle this case appropriately
	}
}
