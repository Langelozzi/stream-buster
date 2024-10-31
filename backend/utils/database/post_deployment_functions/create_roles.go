package post_deployment_functions

import (
	"github.com/STREAM-BUSTER/stream-buster/models"
	"gorm.io/gorm"
)

// InsertRoles inserts predefined roles into the database
func InsertRoles(db *gorm.DB) error {
	// Define the roles
	roles := []models.Role{
		{
			ID:              1,
			Name:            "admin",
			MaxRequestCount: nil, // MaxRequestCount is null
		},
		{
			ID:              2,
			Name:            "basic",
			MaxRequestCount: intPtr(20), // MaxRequestCount is 20
		},
	}

	// Insert roles into the database
	for _, role := range roles {
		if err := db.Create(&role).Error; err != nil {
			return err
		}
	}

	return nil
}

// Helper function to get a pointer to an int
func intPtr(i int) *int {
	return &i
}
