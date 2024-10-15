package models

import "time"

// Document represents a document in the system
type Document struct {
	ID                uint          `gorm:"primaryKey"`
	FileStorageURL    string        `gorm:"not null"`
	UploadedAt        time.Time     `gorm:"default:GETDATE()"`
	CreditValue       int           `gorm:"default:0"`
	FileType          string        `gorm:"default:''"`
	FileName          string        `gorm:"not null"`
	FilePath          string        `gorm:"default:''"`
	Tags              []DocumentTag `gorm:"foreignKey:DocumentID"`
	CategoryID        uint          `gorm:"null;default:null"`
	Category          Category      `gorm:"foreignKey:CategoryID"`
	DeletedAt         *time.Time    `gorm:"index"`
	Address           string        `gorm:"default:''"`
	ApprovedAt        *time.Time    `gorm:"index"`
	ApprovedBy        uint          `gorm:"index;default:null"`
	ApprovedByUser    User          `gorm:"foreignKey:ApprovedBy"`
	AdditionalDetails string        `gorm:"type:text"`
	Description       string        `gorm:"type:text"`
}
