
package models

import "time"

// User represents the user of the system
type MediaType struct {
	ID                 uint   `gorm:"primaryKey"`
	Name string
	Description string
	DeletedAt          *time.Time     `gorm:"index"`
	CreatedAt          *time.Time     `gorm:"index"`
}
