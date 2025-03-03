package post_deployment_functions

import (
	"github.com/STREAM-BUSTER/stream-buster/models/db"
	"gorm.io/gorm"
	"log"
)

func UpdateNullAvailability(database *gorm.DB) error {
	// Update availability to -1 where it's NULL
	result := database.Model(&db.Media{}).Where("availability IS NULL").Update("availability", -1)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		log.Printf("Updated %d media records to set availability = -1", result.RowsAffected)
	} else {
		log.Print("No media records needed availability updates.")
	}

	return nil
}
