package api

import (
	dbModels "github.com/STREAM-BUSTER/stream-buster/models/db"
	"time"
)

type TV struct {
	MediaID uint
	Media   *dbModels.Media

	SeasonCount  int
	EpisodeCount int

	FirstAirDate  *time.Time
	LastAirDate   *time.Time
	BackdropImage string

	Seasons []*Season
}
