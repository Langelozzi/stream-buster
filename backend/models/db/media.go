package db

import (
	"time"
)

type Media struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	TMDBID int  `gorm:"unique"`

	Title        string
	Overview     string
	PosterImage  string
	Availability int `gorm:"default:-1"`

	MediaTypeId uint       `gorm:"foreignKey:ID"`
	MediaType   *MediaType `gorm:"constraint:OnDelete:SET NULL;"`

	DeletedAt *time.Time `gorm:"index"`
	CreatedAt *time.Time `gorm:"index"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"` // Auto-update timestamp on record update

	Genres []*Genre `gorm:"many2many:media_genres;"`
}
