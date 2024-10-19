package database

import (
	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/STREAM-BUSTER/stream-buster/models/db"
	"github.com/STREAM-BUSTER/stream-buster/utils/database/post_deployment_functions"
	"log"
)

func InitializeDb() {
	database := GetInstance()

	// List of all models
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Config{},
		&models.UserConfig{},
		&db.CurrentlyWatching{},
		&db.Media{},
		&db.MediaType{},
		&db.Watchlist{},
	}

	// Automatically migrate schema for each model
	for _, model := range modelsToMigrate {
		if err := database.AutoMigrate(model); err != nil {
			log.Fatalf("Failed to migrate model %T: %v", model, err)
		}
		log.Printf("Successfully migrated model %T", model)
	}

	post_deployment_functions.CreateAdminUser(database)

	log.Print("Database initialized successfully.")
}
