package controllers

import (
	"net/http"
	"strconv"

	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service interfaces.AuthServiceInterface
}

func NewAuthController(service interfaces.AuthServiceInterface) *AuthController {
	return &AuthController{
		service: service,
	}
}

// LoginUser logs in valid users
// @Summary Logs in valid users
// @Description Authenticates a user using the provided username and password, and returns a JWT token in a cookie if successful.
// @Tags Auth
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {string} string "Successfully logged in, JWT set in cookie"
// @Failure 400 {object} map[string]interface{} "Invalid username or password"
// @Router /auth/login [post]
func (contr *AuthController) LoginUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	validCredentials, err := contr.service.CheckCredentials(username, password)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error validating users credentials")
	}
	if validCredentials {

		tokenString, err := contr.service.CreateToken(username)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating token")
			return
		}
		refreshTokenString, err := contr.service.CreateRefreshToken(username)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating refreshToken")
			return
		}
		maxRefreshTokenAge, err := strconv.Atoi(utils.GetEnvVariable("REFRESH_TOKEN_EXPIRATION_TIME"))

		c.SetCookie(
			"refreshToken",                 // Name of the cookie
			refreshTokenString,             // Value of the cookie
			maxRefreshTokenAge,             // MaxAge (7 days)
			"/",                            // Path
			utils.GetEnvVariable("DOMAIN"), // Domain
			false,                          // Secure flag (whether the cookie should be sent only over HTTPS)
			false,                          // HttpOnly flag (whether the cookie is inaccessible to JavaScript)
		)

		c.SetCookie(
			"token",                        // Name of the cookie
			tokenString,                    // Value of the cookie
			3600,                           // MaxAge (1 hour)
			"/",                            // Path
			utils.GetEnvVariable("DOMAIN"), // Domain
			false,                          // Secure flag (whether the cookie should be sent only over HTTPS)
			false,                          // HttpOnly flag (whether the cookie is inaccessible to JavaScript)
		)

		c.String(http.StatusOK, "Autorized")

	} else {
		c.String(http.StatusUnauthorized, "Invalid credentials")
	}
}

func (contr *AuthController) CreateUser(c *gin.Context) {

}
