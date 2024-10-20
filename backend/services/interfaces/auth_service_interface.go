package interfaces

import "github.com/golang-jwt/jwt/v4"

type AuthServiceInterface interface {
	CreateToken(username string) (string, error)
	// authenticateMiddleware(c *gin.Context)
	VerifyToken(tokenString string) (*jwt.Token, error)
}
