package models

import (
	"time"
)

// Role represents a role in the system
type Role struct {
	ID              uint      `gorm:"primaryKey"`
	Name            string    `gorm:"not null;unique"`
	MaxRequestCount *int      `gorm:"default:null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
