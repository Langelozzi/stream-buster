package database

import (
	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/STREAM-BUSTER/stream-buster/models/db"
	"github.com/STREAM-BUSTER/stream-buster/utils/database/post_deployment_functions"
	"gorm.io/gorm"
	"log"
)

func InitializeDb() {
	database := GetInstance()

	// Ensure the database connection is valid
	if database == nil {
		log.Fatal("Database connection is not established.")
		return
	}

	// List of all models
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Config{},
		&models.UserConfig{},
		&models.Role{},
		&models.UserRole{},
		&models.Endpoint{},
		&models.UserEndpointUsage{},
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

	// Post deployment scripts to populate some data in db
	if err := runPostDeploymentScripts(database); err != nil {
		log.Printf("Post-deployment script error: %v", err)
	}

	log.Print("Database initialized successfully.")
}

func runPostDeploymentScripts(database *gorm.DB) error {
	post_deployment_functions.CreateUserTotalRequestCountView(database)

	post_deployment_functions.InsertRoles(database)

	post_deployment_functions.CreateAdminUser(database)

	post_deployment_functions.CreateTestData(database)

	return nil
}
