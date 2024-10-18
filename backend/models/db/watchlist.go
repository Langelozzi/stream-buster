package db

import (
	"github.com/STREAM-BUSTER/stream-buster/models"
	"time"
)

type Watchlist struct {
	UserID uint `gorm:"foreignKey:ID"`
	User   *models.User

	MediaId string `gorm:"foreignKey:ID"`
	Media   *Media `gorm:"constraint:OnDelete:SET NULL;"`

	DeletedAt *time.Time `gorm:"index"`
	CreatedAt *time.Time `gorm:"index"`
}
