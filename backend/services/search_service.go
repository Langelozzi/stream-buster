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

// SearchMultiMedia searches for multimedia content based on a query
func (service *SearchService) SearchMultiMedia(query string) ([]interface{}, error) {
	searchResults, err := service.dao.SearchMultiMedia(query)
	if err != nil {
		return nil, err
	}

	return searchResults, nil

	//var validMedia []interface{}
	//for _, item := range searchResults {
	//	valid := doesMediaContentExist(item)
	//	if valid {
	//		validMedia = append(validMedia, item)
	//	}
	//
	//	// trying to combat rate limiting
	//	time.Sleep(1000 * time.Millisecond)
	//}
	//
	//if len(validMedia) == 0 {
	//	return []interface{}{}, nil
	//}
	//
	//return validMedia, nil

	//var wg sync.WaitGroup
	//results := make([]mediaCheckResult, len(searchResults)) // Pre-allocate slice for results
	//
	//// Start all asynchronous checks concurrently
	//for i, item := range searchResults {
	//	wg.Add(1)
	//	go func(index int, mediaItem interface{}) {
	//		defer wg.Done()
	//		valid := doesMediaContentExist(mediaItem)                        // Call the check function
	//		results[index] = mediaCheckResult{item: mediaItem, valid: valid} // Store result in the pre-allocated slice
	//	}(i, item)
	//}
	//
	//// Wait for all goroutines to finish
	//wg.Wait()
	//
	//// Now that all values have been fetched, iterate through the results slice
	//var validMedia []interface{}
	//for _, result := range results {
	//	if result.valid {
	//		validMedia = append(validMedia, result.item) // Append valid media item
	//	}
	//}
	//
	//if len(validMedia) == 0 {
	//	return []interface{}{}, nil
	//}
	//
	//return validMedia, nil
}

// mediaCheckResult holds the media item and its validity
type mediaCheckResult struct {
	item  interface{}
	valid bool
}

func doesMediaContentExist(item interface{}) bool {
	switch v := item.(type) {
	case *api.TV:
		return utils.DoesContentExist(v.Media.TMDBID, true)
	case *api.Movie:
		return utils.DoesContentExist(v.Media.TMDBID, false)
	default:
		return false
	}
}
