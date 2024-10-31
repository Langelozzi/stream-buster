package interfaces

import (
	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthServiceInterface interface {
	Register(user models.User) (*models.User, error)
	CreateToken(username string) (string, error)
	CreateRefreshToken(username string) (string, error)
	// authenticateMiddleware(c *gin.Context)
	CheckCredentials(password string, user *models.User) bool
	VerifyToken(tokenString string) (*jwt.Token, error)
	RefreshToken(refreshTokenString string) (string, error)
	SetTokenCookie(c *gin.Context, tokenString string)
	HashPassword(password string) (string, error)
}
