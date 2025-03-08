package models

import (
	"time"
)

type Endpoint struct {
	ID        uint64     `gorm:"primaryKey"`
	Method    string     `gorm:"not null: True"`
	Path      string     `gorm:"not null: True"`
	DeletedAt *time.Time `gorm:"index"`
	CreatedAt *time.Time `gorm:"index"`
}
