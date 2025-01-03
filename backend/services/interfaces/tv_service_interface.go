package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models/api"

type TVServiceInterface interface {
	GetTVDetails(id int) (*api.TV, error)
	GetEpisodesInSeason(seriesId int, seasonNum int) ([]*api.Episode, error)
}
