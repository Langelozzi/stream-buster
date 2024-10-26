package api

import (
	dbModels "github.com/STREAM-BUSTER/stream-buster/models/db"
	"time"
)

type Movie struct {
	MediaID uint
	Media   *dbModels.Media

	PosterPath string

	ReleaseDate *time.Time
	Runtime     int
}
