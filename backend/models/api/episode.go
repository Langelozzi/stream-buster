package api

import dbModels "github.com/STREAM-BUSTER/stream-buster/models/db"

type Episode struct {
	MediaID uint
	Media   *dbModels.Media

	Name     string
	Overview string

	EpisodeTMDBID string
	EpisodeNumber int

	StillPath string
	Runtime   int

	SeasonNumber int
}
