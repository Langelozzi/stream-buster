package post_deployment_functions

import (
	"errors"
	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"time"
)

func CreateEndpointRecords(db *gorm.DB, router *gin.Engine) {
	// Retrieve all routes defined in the router
	routes := router.Routes()

	for _, route := range routes {
		// Check if the route already exists in the Endpoint table
		var endpoint models.Endpoint
		result := db.Where("path = ? AND method = ?", route.Path, route.Method).First(&endpoint)

		currentTime := time.Now()
		// If the route doesn't exist, create a new record
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// Insert new route into the Endpoint table
				newEndpoint := models.Endpoint{
					Path:      route.Path,
					Method:    route.Method,
					CreatedAt: &currentTime,
				}

				if err := db.Create(&newEndpoint).Error; err != nil {
					log.Println("Error inserting endpoint:", err)
				} else {
					log.Printf("Inserted new endpoint: %s %s\n", route.Method, route.Path)
				}
			} else {
				log.Println("Error fetching endpoint:", result.Error)
			}
		} else {
			log.Printf("Endpoint already exists: %s %s\n", route.Method, route.Path)
		}
	}
}
