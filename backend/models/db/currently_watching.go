package db

import (
	"github.com/STREAM-BUSTER/stream-buster/models"
	"time"
)

type CurrentlyWatching struct {
	UserID uint         `gorm:"primaryKey"`
	User   *models.User `gorm:"constraint:OnDelete:CASCADE;"`

	MediaId uint   `gorm:"primaryKey"`
	Media   *Media `gorm:"constraint:OnDelete:SET NULL;"`

	EpisodeNumber int
	SeasonNumber  int

	DeletedAt *time.Time `gorm:"index"`
	CreatedAt *time.Time `gorm:"index"`
}
