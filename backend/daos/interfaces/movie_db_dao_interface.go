package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models/api"

type MovieDatabaseDaoInterface interface {
	SearchMultiMedia(query string, page int) (*api.SearchPage, error)
	SearchMovies(query string, page int) (*api.SearchPage, error)
	SearchTV(query string, page int) (*api.SearchPage, error)
	GetTVDetails(id int) (*api.TV, error)
	GetMovieDetails(id int) (*api.Movie, error)
	GetEpisodesInSeason(seriesId int, seasonNum int) ([]*api.Episode, error)
	GetTrendingMovies(timeWindow string, page int) (*api.SearchPage, error)
	GetTrendingTV(timeWindow string, page int) (*api.SearchPage, error)
}
