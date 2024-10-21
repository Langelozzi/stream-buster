package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte(utils.GetEnvVariable(utils.GetEnvVariable("JWT_SECRET")))
var refreshTokenValidLengthSeconds = 60 * 60 * 24 * 6 // 7 days
type AuthService struct {
	Dao iDao.AuthDaoInterface
}

func NewAuthService(dao interfaces.AuthDaoInterface) *AuthService {
	return &AuthService{}
}

// Function to create JWT tokens with claims
func (service AuthService) CreateToken(username string) (string, error) {
	// Create a new JWT token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,                         // Subject (user identifier)
		"iss": "auth-service",                   // Issuer
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
	})

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	// Print information about the created token
	// fmt.Printf("Token claims added: %+v\n", claims)
	return tokenString, nil
}

func (service AuthService) CreateRefreshToken(username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  username,                                                                           // Subject (user identifier)
		"type": "refresh-token",                                                                    // type
		"exp":  time.Now().Add(time.Second * time.Duration(refreshTokenValidLengthSeconds)).Unix(), // Expiration time
		"iat":  time.Now().Unix(),                                                                  // Issued at
	})
	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (service AuthService) VerifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
}

// Function to verify JWT tokens
func (service AuthService) AuthenticateMiddleware(c *gin.Context) {
	// Retrieve the token from the cookie
	tokenString, err := c.Cookie("token")
	if err != nil {
		// fmt.Println("Token missing in cookie")
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	// Verify the token
	token, err := service.VerifyToken(tokenString)
	if err != nil {
		service.handleRefreshToken(c)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		c.Set("user", claims)
	} else {
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	c.Next()
}

func (service AuthService) HandleRefreshToken(refreshTokenString string) (string, error) {
	validToken, err := service.VerifyToken(refreshTokenString)
	if err != nil {
		return "", err
	}

	if !validToken.Valid {
		return "", fmt.Errorf("Refresh Token is invalid")
	}

	claims, ok := validToken.Claims.(jwt.MapClaims)
	if ok {
		// create new token
		username, ok := claims["username"].(string)
		if !ok {

		}
		newAccessToken, err := service.CreateToken(username)
		if err != nil {
			// return some error
			return "", err
		}
		return newAccessToken, nil
	}
	return "", fmt.Errorf("Error parsing claims")
}

func (service AuthService) CheckCredentials(username string, password string) (bool, error) {
	return true, nil
}
