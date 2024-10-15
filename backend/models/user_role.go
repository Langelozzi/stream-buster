package models

import "time"

// UserRole represents the roles assigned to a user
type UserRole struct {
	UserID    uint       `gorm:"primaryKey"`
	RoleID    uint       `gorm:"primaryKey"`
	Role      Role       `gorm:"foreignKey:RoleID"`
	User      User       `gorm:"foreignKey:UserID"`
	DeletedAt *time.Time `gorm:"index"`
}
