package db

import (
	"time"
)

type Media struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	TMDBID int  `gorm:"unique"`

	Title       string
	Overview    string
	PosterImage string

	MediaTypeId uint       `gorm:"foreignKey:ID"`
	MediaType   *MediaType `gorm:"constraint:OnDelete:SET NULL;"`

	DeletedAt *time.Time `gorm:"index"`
	CreatedAt *time.Time `gorm:"index"`

	Genres []*Genre `gorm:"many2many:media_genres;"`
}
