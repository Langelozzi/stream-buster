
package models

import "time"

// User represents the user of the system
type Media struct {
	ID                 uint   `gorm:"primaryKey"`
	TMDBID          int 
	Title string
	PosterImage string
	DeletedAt          *time.Time     `gorm:"index"`
	CreatedAt          *time.Time     `gorm:"index"`
}
