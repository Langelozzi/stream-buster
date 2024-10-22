package middlewares

import (
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func Auth(service interfaces.AuthServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No valid refresh token"})
			c.Abort()
			return
		}

		token, err := service.VerifyToken(tokenString)
		if err != nil || !token.Valid {
			refreshTokenString, err := c.Cookie("refreshToken")
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "No valid refresh token"})
				c.Abort()
				return
			}

			accessTokenString, err := service.RefreshToken(refreshTokenString)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to refresh token"})
				c.Abort()
				return
			}

			service.SetTokenCookie(c, accessTokenString)

			token, err = service.VerifyToken(accessTokenString)
			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refreshed token"})
				c.Abort()
				return
			}
		}

		// Extract claims from the verified token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user", claims)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}
