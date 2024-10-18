
package models

import "time"

type Season struct {
	MediaId           string 
	Media             Media

	Name string
	Overview string

	SeasonTMDBID string
	SeasonNumber int

	PosterPath string
	Episodes Episode[]
}
