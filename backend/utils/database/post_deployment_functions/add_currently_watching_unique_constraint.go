package post_deployment_functions

import (
	dbModels "github.com/STREAM-BUSTER/stream-buster/models/db"
	"gorm.io/gorm"
	"log"
)

func AddUniqueConstraintToCurrentlyWatching(db *gorm.DB) error {
	// Define the unique constraint name
	constraintName := "unique_user_media"

	// Ensure the table exists before altering it
	if !db.Migrator().HasTable(&dbModels.CurrentlyWatching{}) {
		log.Printf("Table 'currently_watchings' does not exist.")
		return nil
	}

	// Add a unique constraint to the 'user_id' and 'media_id' columns
	err := db.Exec("ALTER TABLE currently_watchings ADD CONSTRAINT " + constraintName + " UNIQUE (user_id, media_id);").Error
	if err != nil {
		log.Printf("Error adding unique constraint: %v", err)
		return err
	}

	log.Printf("Successfully added unique constraint %s to 'user_id' and 'media_id'.", constraintName)
	return nil
}
