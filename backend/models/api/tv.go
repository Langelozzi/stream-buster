package models

import "time"

// User represents the user of the system
type TV struct {
	MediaId           string 
	Media             Media

	Overview string

	SeasonCount int
	EpisodeCount int

	Seasons Season[]
	FirstAirDate date
}
