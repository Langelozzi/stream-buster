package api

import dbModels "github.com/STREAM-BUSTER/stream-buster/models/db"

type Season struct {
	MediaID uint
	Media   *dbModels.Media

	SeasonTMDBID string
	SeasonNumber int
	EpisodeCount int

	Name     string
	Overview string

	PosterPath string
	Episodes   []*Episode
}
