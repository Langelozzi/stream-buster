package database

import (
	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/STREAM-BUSTER/stream-buster/utils/database/post_deployment_functions"
	"log"
)

func InitializeDb() {
	db := GetInstance()

	// List of all models
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Config{},
		&models.UserConfig{},
	}

	// Automatically migrate schema for each model
	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("Failed to migrate model %T: %v", model, err)
		}
		log.Printf("Successfully migrated model %T", model)
	}

	post_deployment_functions.CreateAdminUser(db)

	log.Print("Database initialized successfully.")
}
