package post_deployment_functions

import (
	"errors"
	"fmt"
	"github.com/STREAM-BUSTER/stream-buster/models"
	"gorm.io/gorm"
)

func CreateAdminUser(db *gorm.DB) {
	// Check if a user with ID 1 already exists
	var existingUser models.User
	result := db.First(&existingUser, 1)
	if result.Error == nil {
		fmt.Println("User with ID 1 already exists. Skipping insert.")
		return
	} else if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Handle other potential errors
		panic(result.Error)
	}

	// Allow identity insert for the user table
	db.Exec("SET IDENTITY_INSERT users ON")

	// Create the admin user
	user := models.User{
		ID:        1,
		FirstName: "Admin",
		LastName:  "LnameAdmin",
		Username:  "admin",
		Password:  "$tre@mBuster",
		DeletedAt: nil,
	}
	db.Create(&user)

	// Disallow identity insert for the user table
	db.Exec("SET IDENTITY_INSERT users OFF")
}
