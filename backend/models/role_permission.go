package models

import "time"

// RolePermission represents the permissions assigned to a role
type RolePermission struct {
	RoleID       uint       `gorm:"primaryKey"`
	PermissionID uint       `gorm:"primaryKey"`
	Role         Role       `gorm:"foreignKey:RoleID"`
	Permission   Permission `gorm:"foreignKey:PermissionID"`
	DeletedAt    *time.Time `gorm:"index"`
}
