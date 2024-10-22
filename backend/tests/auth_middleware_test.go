package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/STREAM-BUSTER/stream-buster/daos"
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/services"
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/auth/test", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	var dao iDao.AuthDaoInterface = daos.NewAuthDao()
	var authService interfaces.AuthServiceInterface = services.NewAuthService(dao)

	accessTokenString, err := authService.CreateToken("username")

	cookie := &http.Cookie{
		Name:     "token",
		Value:    accessTokenString,
		MaxAge:   3600,  // 1 hour
		HttpOnly: false, // Secure the cookie by not allowing JavaScript access
	}
	req.AddCookie(cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status 200 OK")
}

func TestMiddleware_refreshToken(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/auth/test", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	var dao iDao.AuthDaoInterface = daos.NewAuthDao()
	var authService interfaces.AuthServiceInterface = services.NewAuthService(dao)

	accessTokenString, err := authService.CreateToken("username")

	cookie := &http.Cookie{
		Name:     "token",
		Value:    accessTokenString,
		MaxAge:   3600,  // 1 hour
		HttpOnly: false, // Secure the cookie by not allowing JavaScript access
	}

	refreshTokenString, err := authService.CreateRefreshToken("username")
	refreshCookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshTokenString,
		MaxAge:   3600,  // 1 hour
		HttpOnly: false, // Secure the cookie by not allowing JavaScript access
	}
	req.AddCookie(cookie)
	req.AddCookie(refreshCookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status 200 OK")
}
