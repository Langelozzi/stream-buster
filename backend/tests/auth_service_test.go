package tests

import (
	"fmt"
	"testing"

	"github.com/STREAM-BUSTER/stream-buster/daos"
	daoInterfaces "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/services"
	servInterfaces "github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/utils"
	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("my_secret_key")

// TestCreateToken_ReturnsToken tests if the function returns a non-empty token string.
func TestCreateToken_ReturnsToken(t *testing.T) {
	var dao daoInterfaces.AuthDaoInterface = daos.NewAuthDao()
	var authService servInterfaces.AuthServiceInterface = services.NewAuthService(dao)

	token, err := authService.CreateToken("testuser")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if token == "" {
		t.Errorf("Expected a token string, got an empty string")
	}
}

// TestCreateToken_ValidToken tests if the generated token is valid and contains the correct claims.
func TestCreateToken_ValidToken(t *testing.T) {
	var dao daoInterfaces.AuthDaoInterface = daos.NewAuthDao()
	var authService servInterfaces.AuthServiceInterface = services.NewAuthService(dao)
	username := "testuser"
	tokenString, err := authService.CreateToken(username)
	fmt.Println(tokenString)
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
	var dao daoInterfaces.AuthDaoInterface = daos.NewAuthDao()
	var authService servInterfaces.AuthServiceInterface = services.NewAuthService(dao)

	token, err := authService.CreateToken("cameron")
	if err != nil {
		t.Fatalf("Error creating token: %v", err)
	}

	verifiedToken, err := authService.VerifyToken(token)
	if err != nil {
		t.Fatalf("Error verifying token %v", err)
	}

	fmt.Println(verifiedToken)

}
