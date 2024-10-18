package db

import "time"

type MediaType struct {
	ID uint `gorm:"primaryKey"`

	Name        string
	Description string

	DeletedAt *time.Time `gorm:"index"`
	CreatedAt *time.Time `gorm:"index"`
}
