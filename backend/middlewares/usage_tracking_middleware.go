package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// UsageTrackingMiddleware tracks the number of requests made by the user
func UsageTrackingMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user info from the context
		userClaims, exists := c.Get("user") // Adjust this according to how you set the user context
		if !exists {
			c.Next() // User not authenticated, just proceed
			return
		}

		userID := uint(userClaims.(jwt.MapClaims)["id"].(float64)) // Assuming "id" is part of the claims

		fmt.Println(userID)
		// Increment the request count for the user
		//var usage models.Usage
		//if err := db.First(&usage, "user_id = ?", userID).Error; err == nil {
		//	usage.RequestCount++
		//	db.Save(&usage)
		//} else {
		//	// Handle the case where Usage record does not exist (if needed)
		//	// This can happen if Usage record is not created for a user yet
		//}

		c.Next() // Call the next handler
	}
}
