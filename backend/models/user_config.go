package models

import "time"

// UserConfig represents configurations specific to a user
type UserConfig struct {
	UserID    uint       `gorm:"primaryKey"`
	ConfigID  uint       `gorm:"primaryKey"`
	User      User       `gorm:"foreignKey:UserID"`
	Config    Config     `gorm:"foreignKey:ConfigID"`
	DeletedAt *time.Time `gorm:"index"`
	CreatedAt *time.Time `gorm:"index"`
}
