package models

import "time"

type Movie struct {
	MediaId           string 
	Media             Media

	Overview string

	PosterPath string

	Genres Genre []
	ReleaseDate date 
	Runtime int
}
