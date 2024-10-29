package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models/api"

type MovieDatabaseDaoInterface interface {
	SearchMultiMedia(query string) ([]interface{}, error)
	GetTVDetails(id int) (*api.TV, error)
}
