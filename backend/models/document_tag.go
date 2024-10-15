package models

import "time"

// DocumentTag represents the relationship between a document and tags
type DocumentTag struct {
	DocumentID uint       `gorm:"primaryKey"`
	TagID      uint       `gorm:"primaryKey"`
	Document   Document   `gorm:"foreignKey:DocumentID"`
	Tag        Tag        `gorm:"foreignKey:TagID"`
	DeletedAt  *time.Time `gorm:"index"`
}
