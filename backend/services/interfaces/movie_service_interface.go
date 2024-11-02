package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models/api"

type MovieServiceInterface interface {
	GetMovieDetails(id int) (*api.Movie, error)
}
