package controllers

import (
	"net/http"
	"strconv"

	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service interfaces.AuthServiceInterface
}

func NewAuthController(service interfaces.AuthServiceInterface) *AuthController {
	return &AuthController{
		Service: service,
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

	validCredentials, err := contr.Service.CheckCredentials(username, password)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error validating users credentials")
	}
	if validCredentials {

		tokenString, err := contr.Service.CreateToken(username)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating token")
			return
		}
		refreshTokenString, err := contr.Service.CreateRefreshToken(username)
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

		contr.Service.SetTokenCookie(c, tokenString)

		c.String(http.StatusOK, "Autorized")

	} else {
		c.String(http.StatusUnauthorized, "Invalid credentials")
	}
}

func (contr *AuthController) CreateUser(c *gin.Context) {

}

func (contr *AuthController) TestAuthMiddleware(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "No user found"})
	}
	c.JSON(200, user)
}
