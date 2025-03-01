package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models/api"

type SearchServiceInterface interface {
	SearchMultiMedia(query string, page int) (*api.SearchPage, error) // []interface{} allows mixed type slice
	//SearchTV(query string) ([]*apiModels.TV, error)
	//SearchMovie(query string) ([]*apiModels.Movie, error)
}
