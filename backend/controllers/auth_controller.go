package controllers

import (
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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

		// fmt.Printf("Token created: %s\n", tokenString)
		c.SetCookie("token", tokenString, 3600, "/", utils.GetEnvVariable("DOMAIN"), false, true)
		c.String(http.StatusOK, "Autorized")
	} else {
		c.String(http.StatusUnauthorized, "Invalid credentials")
	}
}

func (contr *AuthController) CreateUser(c *gin.Context) {

}
