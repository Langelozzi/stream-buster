package database

import (
	"github.com/DOC-SWAP/Docswap-backend/utils/database/post_deployment_functions"
	"log"

	"github.com/DOC-SWAP/Docswap-backend/models"
)

func InitializeDb() {
	db := GetInstance()

	// List of all models
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Document{},
		&models.Config{},
		&models.Role{},
		&models.Tag{},
		&models.Permission{},
		&models.UserDocument{},
		&models.DocumentTag{},
		&models.UserConfig{},
		&models.UserRole{},
		&models.RolePermission{},
		&models.Category{},
		&models.UserTag{},
	}

	// Automatically migrate schema for each model
	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("Failed to migrate model %T: %v", model, err)
		}
		log.Printf("Successfully migrated model %T", model)
	}

	post_deployment_functions.CreateCategories(db)
	post_deployment_functions.CreateAdminUser(db)

	log.Print("Database initialized successfully.")
}
