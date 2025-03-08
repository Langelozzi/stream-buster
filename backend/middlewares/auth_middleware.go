package middlewares

import (
	"net/http"

	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(service interfaces.AuthServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")

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
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to refresh token" + err.Error()})
				c.Abort()
				return
			}

			service.SetCookie(c, "token", accessTokenString, 3600)

			token, err = service.VerifyToken(accessTokenString)
			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refreshed token"})
				c.Abort()
				return
			}
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("ID", claims["id"])
			c.Set("Email", claims["email"])
			c.Set("FirstName", claims["fname"])
			c.Set("LastName", claims["lname"])
			c.Set("Issuer", claims["iss"])
			c.Set("Exp", claims["exp"])
			c.Set("Iat", claims["iat"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}
		c.Next()
	}
}
