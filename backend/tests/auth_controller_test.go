package test

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	// Create the form data for the request
	form := url.Values{}
	form.Add("username", "testuser")
	form.Add("password", "testpassword")

	// Create the request, set the content-type to form-encoded
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/auth/login", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status 200 OK")

	// Optionally, check for the cookie or body content
	cookies := resp.Cookies()
	assert.NotEmpty(t, cookies, "Expected JWT cookie in response")
	for _, cookie := range cookies {
		if cookie.Name == "token" {
			assert.NotEmpty(t, cookie.Value, "Expected token value in cookie")
		}
	}
}
