package api

import dbModels "github.com/STREAM-BUSTER/stream-buster/models/db"

type Season struct {
	MediaID uint
	Media   *dbModels.Media

	Name     string
	Overview string

	SeasonTMDBID string
	SeasonNumber int

	PosterPath string
	Episodes   []*Episode
}
