package api

import (
	dbModels "github.com/STREAM-BUSTER/stream-buster/models/db"
	"time"
)

type TV struct {
	MediaID uint
	Media   *dbModels.Media

	Overview string

	SeasonCount  int
	EpisodeCount int

	Seasons      []*Season
	FirstAirDate *time.Time
}
