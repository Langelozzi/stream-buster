package post_deployment_functions

import (
	"errors"
	"fmt"
	"github.com/STREAM-BUSTER/stream-buster/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateAdminUser(db *gorm.DB) error {
	// Check if a user with ID 1 already exists
	var existingUser models.User
	result := db.First(&existingUser, 1)
	if result.Error == nil {
		fmt.Println("User with ID 1 already exists. Skipping insert.")
		return result.Error
	} else if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Handle other potential errors
		panic(result.Error)
	}

	unHashedPassword := "111"
	// Generate a hashed password with bcrypt using a cost of bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(unHashedPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("error hashing password:", err)
		return err
	}

	password := string(hashedPassword)
	// Create the admin user
	user := models.User{
		ID:        1,
		Email:     "admin@admin.com",
		FirstName: "Admin",
		LastName:  "LnameAdmin",
		Password:  password,
		DeletedAt: nil,
	}

	// Create usage record
	usage := models.Usage{
		UserID:       user.ID,
		RequestCount: 0,
	}

	// Use a transaction to ensure both user and userRole are inserted atomically
	err = db.Transaction(func(tx *gorm.DB) error {
		// Create the user
		if err := tx.Create(&user).Error; err != nil {
			return err // Rollback if user creation fails
		}

		// Create the userRole record
		userRole := models.UserRole{
			UserID: user.ID,
			RoleID: 1, // Assuming the role ID for admin is 1
		}

		if err := tx.Create(&userRole).Error; err != nil {
			return err // Rollback if userRole creation fails
		}

		// Create the usage record
		if err := tx.Create(&usage).Error; err != nil {
			return err // Rollback if usage creation fails
		}

		return nil // Commit the transaction
	})

	if err != nil {
		fmt.Println("Error creating admin user and related records:", err)
	}

	return nil
}
