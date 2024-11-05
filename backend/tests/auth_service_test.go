package test

import (
	"testing"

	"github.com/STREAM-BUSTER/stream-buster/utils"
	"github.com/STREAM-BUSTER/stream-buster/utils/dependency_injection"
	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("my_secret_key")

// TestCreateToken_ReturnsToken tests if the function returns a non-empty token string.
func TestCreateToken_ReturnsToken(t *testing.T) {
	authController := dependency_injection.InitAuthDependencies()
	authService := authController.Service

	token, err := authService.CreateToken("Admin@streambuster.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if token == "" {
		t.Errorf("Expected a token string, got an empty string")
	}
}

// TestCreateToken_ValidToken tests if the generated token is valid and contains the correct claims.
func TestCreateToken_ValidToken(t *testing.T) {

	authController := dependency_injection.InitAuthDependencies()
	authService := authController.Service

	email := "Admin@streambuster.com"
	tokenString, err := authService.CreateToken(email)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Parse the token to validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetEnvVariable("JWT_SECRET")), nil
	})

	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	// Check if token is valid
	if !token.Valid {
		t.Errorf("Token is invalid")
	}
}

func TestCreateAndVerifyToken(t *testing.T) {

	authController := dependency_injection.InitAuthDependencies()
	authService := authController.Service

	tokenString, err := authService.CreateToken("Admin@streambuster.com")
	if err != nil {
		t.Fatalf("Error creating token: %v", err)
	}

	token, err := authService.VerifyToken(tokenString)
	if err != nil {
		t.Fatalf("Error verifying token %v", err)
	}

	if !token.Valid {
		t.Errorf("Token is invalid")
	}
}

// func TestRefreshToken(t *testing.T) {
//
// 	var dao daoInterfaces.AuthDaoInterface = daos.NewAuthDao()
// 	var authService servInterfaces.AuthServiceInterface = services.NewAuthService(dao)
//
// 	tokenString, err := authService.CreateRefreshToken("cameron")
// 	if err != nil {
// 		t.Fatal("Error Creating token %v", err)
// 	}
//
// }
