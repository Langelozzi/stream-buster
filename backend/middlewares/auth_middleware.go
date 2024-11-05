package middlewares

import (
	"net/http"

	"github.com/STREAM-BUSTER/stream-buster/models/auth"
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(service interfaces.AuthServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil || tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No valid access token"})
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
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Manually map claims to the UserClaims struct
			userClaims := auth.UserClaims{
				ID:        int(claims["id"].(float64)),
				Email:     claims["email"].(string),
				FirstName: claims["fname"].(string),
				LastName:  claims["lname"].(string),
				Issuer:    claims["iss"].(string),
				Exp:       int64(claims["exp"].(float64)), // JWT encodes numbers as float64
				Iat:       int64(claims["iat"].(float64)),
			}

			// Store the user claims in the context
			c.Set("user", userClaims)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}
		c.Next()
	}
}
