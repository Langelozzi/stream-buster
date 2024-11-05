package db

import (
	"github.com/STREAM-BUSTER/stream-buster/models"
	"time"
)

type CurrentlyWatching struct {
	UserID uint         `gorm:"foreignKey:ID"`
	User   *models.User `gorm:"constraint:OnDelete:CASCADE;"`

	MediaId uint   `gorm:"foreignKey:ID"`
	Media   *Media `gorm:"constraint:OnDelete:SET NULL;"`

	EpisodeNumber int
	SeasonNumber  int

	DeletedAt *time.Time `gorm:"index"` // Use pointer to allow null
	CreatedAt *time.Time `gorm:"index"` // Use pointer to allow null
}
