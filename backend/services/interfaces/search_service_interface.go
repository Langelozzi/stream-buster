package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models/api"

type SearchServiceInterface interface {
	SearchMultiMedia(query string, page int) (*api.SearchPage, error)
	SearchMovies(query string, page int) (*api.SearchPage, error)
	SearchTV(query string, page int) (*api.SearchPage, error)
	SearchTrendingMovies(timeWindow string, page int) (*api.SearchPage, error)
	SearchTrendingTV(timeWindow string, page int) (*api.SearchPage, error)
}
