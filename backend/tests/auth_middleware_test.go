package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMiddleware(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/auth/test", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjk2MzQ2NDIsImlhdCI6MTcyOTYzMTA0MiwiaXNzIjoiYXV0aC1zZXJ2aWNlIiwic3ViIjoidGVzdHVzZXIifQ.XRdajlVpkYQQeWaLwt3ZMzYl7IXFaMavQ522JRmzmxs",
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

	cookie := &http.Cookie{
		Name:     "token",
		Value:    "eyJhbGcOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjk2MjM5MDUsImlhdCI6MTcyOTYyMDMwNSwiaXNzIjoiYXV0aC1zZXJ2aWNlIiwic3ViIjoidGVzdHVzZXIifQ.BZc-7UlEvV5oj0P2mfxVYHgXGJjCoH3uDne-3i7e35I",
		MaxAge:   3600,  // 1 hour
		HttpOnly: false, // Secure the cookie by not allowing JavaScript access
	}

	refreshCookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAyMjcwOTMsImlhdCI6MTcyOTYyMjI5Mywic3ViIjoidGVzdHVzZXIiLCJ0eXBlIjoicmVmcmVzaC10b2tlbiJ9.pohYGI4c9abWg3lo13S4YzHpQgYKSJVNXA0jRnVUiXc",
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
