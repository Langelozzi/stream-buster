package models

import "time"

type Episode struct {
	MediaId           string 
	Media             Media

	Name string
	Overview string

	EpisodeTMDBID string
	EpisodeNumber int

	StillPath string
	Runtime int 

	SeasonNumber int 
}
