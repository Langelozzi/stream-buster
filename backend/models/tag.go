package models

import "time"

// Tag represents a tag that can be applied to documents
type Tag struct {
	ID          uint          `gorm:"primaryKey"`
	Name        string        `gorm:"unique;not null"`
	Description string        `gorm:"default:''"`
	Documents   []DocumentTag `gorm:"foreignKey:TagID"`
	DeletedAt   *time.Time    `gorm:"index"`
}
