package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models/api"

type SearchServiceInterface interface {
	SearchMultiMedia(query string, page int) (*api.SearchPage, error)
	SearchTrendingMovies(timeWindow string, page int) (*api.SearchPage, error)
}
