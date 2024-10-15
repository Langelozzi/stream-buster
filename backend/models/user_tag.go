package models

import "time"

// UserTag represents the association between a user and a tag
type UserTag struct {
	UserID    uint       `gorm:"not null;index"`
	TagID     uint       `gorm:"not null;index"`
	DeletedAt *time.Time `gorm:"index"`
	User      User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tag       Tag        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
