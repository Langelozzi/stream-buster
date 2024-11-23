package middlewares

import (
	"net/http"

	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/STREAM-BUSTER/stream-buster/utils/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UsageTrackingMiddleware tracks the number of requests made by the user
func UsageTrackingMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user info from the context

		user, err := auth.GetUserFromContext(c)
		if err != nil {
			c.Redirect(http.StatusUnauthorized, "/login")
		}

		userID := uint(user.ID) // Assuming "id" is part of the claims

		// Increment the request count for the user
		var usage models.Usage
		if err := db.First(&usage, "user_id = ?", userID).Error; err == nil {
			usage.RequestCount++
			db.Save(&usage)
		} else {
			// Handle the case where Usage record does not exist (if needed)
			// This can happen if Usage record is not created for a user yet
		}

		c.Next() // Call the next handler
	}
}
