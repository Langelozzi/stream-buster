package models

import "time"

// Permission represents a system permission
type Permission struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Roles       []RolePermission `gorm:"foreignKey:PermissionID"`
	DeletedAt   *time.Time       `gorm:"index"`
}
