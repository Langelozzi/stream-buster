package models

type Category struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	ParentID    *uint     `gorm:"index;default:null"`
	Parent      *Category `gorm:"foreignKey:ParentID"`
}
