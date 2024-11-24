package middlewares

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/STREAM-BUSTER/stream-buster/utils/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UsageTrackingMiddleware tracks the number of requests made by the user
func UsageTrackingMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the endpoint path and userid
		endpointPath := c.FullPath()

		user, err := auth.GetUserFromContext(c)
		if err != nil {
			c.Redirect(http.StatusUnauthorized, "/login")
		}

		userID := uint(user.ID) // Assuming "id" is part of the claims

		var endpoint models.Endpoint
		if err := db.Where("path = ?", endpointPath).First(&endpoint).Error; err != nil {
			// Log or handle error if endpoint is not found in the Endpoint table
			c.Next()
			return
		}

		// Check if there's already a usage record for this user and endpoint
		var usage models.UserEndpointUsage
		err = db.Where("user_id = ? AND endpoint_id = ?", userID, endpoint.ID).First(&usage).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Record does not exist, create it
			usage = models.UserEndpointUsage{
				UserID:       uint(userID), // Cast userID to appropriate type
				EndpointID:   uint(endpoint.ID),
				RequestCount: 1,
			}
			if err := db.Create(&usage).Error; err != nil {
				log.Println("Error creating usage record:", err)
			}
		} else {
			// Record exists, increment the request count and update last_access
			db.Model(&usage).Updates(map[string]interface{}{
				"request_count": gorm.Expr("request_count + ?", 1),
				"last_access":   time.Now(),
			})
		}

		c.Next()
	}
}
