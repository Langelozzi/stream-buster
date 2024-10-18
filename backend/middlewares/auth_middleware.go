package middlewares

import (
	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//authHandler := dependency_injection.InitAuthDependencies()
		//userService := dependency_injection.InitUserServiceDependencies()

		// Get the Authentication Header
		authHeader := c.GetHeader("Authorization")

		// Check if the token is missing
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Missing Authorization Header"})
			c.Abort()
			return
		}

		// Remove the Bearer prefix
		authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

		//// Validate the token
		//token, err := authHandler.ParseAndValidateToken(authHeader)
		//if err != nil {
		//	c.JSON(401, gin.H{"error": err.Error()})
		//	c.Abort()
		//	return
		//}
		//
		//// If the token is valid, get the external userId from claims
		//claims := token.Claims.(*auth.TokenClaims)
		//externalUserId := claims.Oid

		// Get the user from the database
		//user, err := userService.GetUserByExternalId(externalUserId, false, false)
		//if err != nil {
		//	c.JSON(500, gin.H{"error": err.Error()})
		//	c.Abort()
		//	return
		//}
		user := models.User{}

		// Set the user in the context
		c.Set("user", user)

		c.Next()
	}
}
