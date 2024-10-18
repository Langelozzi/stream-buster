package models

import "time"

// User represents the user of the system
type CurrentlyWatching struct {
	UserID                 uint   `gorm:"foreignKey:UserID"`
	User               User 

	MediaId           string `gorm:"foreignKey:MediaID"`
	Media             Media

	MediaTypeId       uint `gorm:"foreignKeyMediaTypeID`
	MediaType         MediaType

	EpisodeNumber     int 
	SeasonNumber     int

	DeletedAt          *time.Time     `gorm:"index"`
	CreatedAt          *time.Time     `gorm:"index"`
}
