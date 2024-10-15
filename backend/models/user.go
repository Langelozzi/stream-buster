package models

import "time"

// User represents the user of the system
type User struct {
	ID                 uint   `gorm:"primaryKey"`
	ExternalUserID     string `gorm:"unique"`
	Username           string `gorm:"unique"`
	AuthenticationType string
	FirstName          string
	LastName           string
	Email              string
	Biography          string
	Documents          []UserDocument `gorm:"foreignKey:UserID"`
	Roles              []UserRole     `gorm:"foreignKey:UserID"`
	Configs            []UserConfig   `gorm:"foreignKey:UserID"`
	DeletedAt          *time.Time     `gorm:"index"`
}
