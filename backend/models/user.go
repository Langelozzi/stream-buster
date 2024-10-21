package models

import "time"

// User represents the user of the system
type User struct {
	ID        uint         `gorm:"primaryKey"`
	Username  string       `gorm:"unique"`
	FirstName string       `gorm:"not null: True"`
	LastName  string       `gorm:"not null: True"`
	Email     string       `gorm:"not null: True"`
	Password  string       `gorm:"not null: True"`
	Configs   []UserConfig `gorm:"foreignKey:UserID"`
	DeletedAt *time.Time   `gorm:"index"`
	CreatedAt *time.Time   `gorm:"index"`
}
