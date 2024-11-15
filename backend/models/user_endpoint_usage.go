package models

import (
	"time"
)

type UserEndpointUsage struct {
	ID           uint      `gorm:"primaryKey"`
	UserID       uint      `gorm:"not null"`
	EndpointID   uint      `gorm:"not null"`
	RequestCount int       `gorm:"default:0"`
	LastAccess   time.Time `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	User     User     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Endpoint Endpoint `gorm:"foreignKey:EndpointID;constraint:OnDelete:CASCADE"`
}
