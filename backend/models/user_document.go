package models

import "time"

// UserDocument represents the relationship between a user and documents
type UserDocument struct {
	UserID     uint `gorm:"primaryKey"`
	DocumentID uint `gorm:"primaryKey"`
	IsOwner    bool
	User       User       `gorm:"foreignKey:UserID"`
	Document   Document   `gorm:"foreignKey:DocumentID"`
	DeletedAt  *time.Time `gorm:"index"`
}
