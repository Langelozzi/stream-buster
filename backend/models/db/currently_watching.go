package db

import (
	"time"

	"github.com/STREAM-BUSTER/stream-buster/models"
)

type CurrentlyWatching struct {
	// Composite primary keys
	UserID uint         `gorm:"primaryKey;foreignKey:ID"`
	User   *models.User `gorm:"constraint:OnDelete:CASCADE;"`

	MediaID uint   `gorm:"primaryKey;foreignKey:ID"`
	Media   *Media `gorm:"constraint:OnDelete:SET NULL;"`

	EpisodeNumber int
	SeasonNumber  int

	DeletedAt *time.Time `gorm:"index"` // Use pointer to allow null
	CreatedAt *time.Time `gorm:"index"` // Use pointer to allow null
}
