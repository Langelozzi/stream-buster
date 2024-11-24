package models

import (
	"time"
)

type Usage struct {
	ID           uint      `gorm:"primaryKey"`
	UserID       uint      `gorm:"unique;not null"` // Ensure uniqueness for one-to-one relationship
	RequestCount int       `gorm:"not null;default:0"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	DeletedAt    time.Time `gorm:"index"`
}
