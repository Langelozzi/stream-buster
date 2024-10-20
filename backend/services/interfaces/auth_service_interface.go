package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthServiceInterface interface {
	createToken(username string) (string, error)
	getRole(username string) string
	authenticateMiddleware(c *gin.Context)
	verifyToken(tokenString string) (*jwt.Token, error)
}
