package services

import (
	"fmt"
	"strconv"
	"time"

	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/models"
	iServices "github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var superSecretKey = []byte(utils.GetEnvVariable(utils.GetEnvVariable("JWT_SECRET")))

type AuthService struct {
	Dao         iDao.AuthDaoInterface
	userService iServices.UserServiceInterface
}

func NewAuthService(dao iDao.AuthDaoInterface, userService iServices.UserServiceInterface) *AuthService {
	return &AuthService{dao, userService}
}

// Function to create JWT tokens with claims
func (service AuthService) CreateToken(email string) (string, error) {
	user, err := service.userService.GetUserByEmail(email, false, false)
	if err != nil {
		return "", err
	}
	// Create a new JWT token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"fname": user.FirstName,
		"lname": user.LastName,
		"iss":   "auth-service",                   // Issuer
		"exp":   time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat":   time.Now().Unix(),                // Issued at
	})

	tokenString, err := claims.SignedString(superSecretKey)
	if err != nil {
		return "", err
	}

	// Print information about the created token
	// fmt.Printf("Token claims added: %+v\n", claims)
	return tokenString, nil
}

func (service AuthService) CreateRefreshToken(email string) (string, error) {
	maxRefreshTokenAge, err := strconv.Atoi(utils.GetEnvVariable("REFRESH_TOKEN_EXPIRATION_TIME"))
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"type":  "refresh-token",
		"exp":   time.Now().Add(time.Second * time.Duration(maxRefreshTokenAge)).Unix(), // Expiration time
		"iat":   time.Now().Unix(),                                                      // Issued at
	})
	tokenString, err := claims.SignedString(superSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (service AuthService) VerifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return superSecretKey, nil
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

func (service AuthService) RefreshToken(refreshTokenString string) (string, error) {
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
		email, ok := claims["sub"].(string)
		if !ok {
			return "", fmt.Errorf("Error parsing username from claims")
		}
		newAccessToken, err := service.CreateToken(email)
		if err != nil {
			// return some error
			return "", err
		}
		return newAccessToken, nil

	} else {

		return "", fmt.Errorf("Error parsing claims")

	}
}

func (service AuthService) SetTokenCookie(c *gin.Context, tokenString string) {
	c.SetCookie(
		"token",                        // Name of the cookie
		tokenString,                    // Value of the cookie
		3600,                           // MaxAge (1 hour)
		"/",                            // Path
		utils.GetEnvVariable("DOMAIN"), // Domain
		false,                          // Secure flag (whether the cookie should be sent only over HTTPS)
		false,                          // HttpOnly flag (whether the cookie is inaccessible to JavaScript)
	)
}

func (service AuthService) CreateUser(firstName, lastName, email, password string) (*models.User, error) {
	// hash the password
	hashPasswordString, err := service.HashPassword(password)
	if err != nil {
	}

	// Create the user object
	newUser := models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hashPasswordString,
	}
	createdUser, err := service.userService.CreateUser(&newUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// HashPassword hashes a raw password string using bcrypt
func (service AuthService) HashPassword(password string) (string, error) {
	// Generate a hashed password with bcrypt using a cost of bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (service AuthService) CheckCredentials(password string, user *models.User) bool {
	return service.CheckPasswordHash(password, user.Password)
}

// CheckPasswordHash compares a hashed password with its possible plain-text equivalent
func (service AuthService) CheckPasswordHash(password, hash string) bool {
	// Compare the password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
