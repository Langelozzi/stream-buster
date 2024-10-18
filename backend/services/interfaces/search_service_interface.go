package interfaces

type SearchServiceInterface interface {
	SearchMultiMedia(query string) ([]interface{}, error) // []interface{} allows mixed type slice
	//SearchTV(query string) ([]*apiModels.TV, error)
	//SearchMovie(query string) ([]*apiModels.Movie, error)
}
