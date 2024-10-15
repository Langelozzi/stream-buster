package models

import "time"

// Config represents a system configuration
type Config struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"not null"`
	Value     string     `gorm:"default:''"`
	DeletedAt *time.Time `gorm:"index"`
}
