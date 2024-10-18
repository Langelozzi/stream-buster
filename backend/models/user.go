package models

import "time"

// User represents the user of the system
type User struct {
	ID                 uint   `gorm:"primaryKey"`
	Username           string `gorm:"unique"`
	FirstName          string
	LastName           string
	Email              string
	Password           string
	Configs            []UserConfig   `gorm:"foreignKey:UserID"`
	DeletedAt          *time.Time     `gorm:"index"`
	CreatedAt          *time.Time     `gorm:"index"`
}
