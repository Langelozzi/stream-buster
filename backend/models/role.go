package models

import "time"

// Role represents a user role in the system
type Role struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Users       []UserRole       `gorm:"foreignKey:RoleID"`
	Permissions []RolePermission `gorm:"foreignKey:RoleID"`
	DeletedAt   *time.Time       `gorm:"index"`
}
