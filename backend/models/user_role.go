package models

// UserRole represents the association between a user and a role
type UserRole struct {
	UserID uint `gorm:"not null;index"`
	RoleID uint `gorm:"not null;index"`

	// Optional associations for eager loading
	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
