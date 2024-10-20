package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/STREAM-BUSTER/stream-buster/models"
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
// @Summary Logs in valid users j
// @Description checks the credentails of the user and returns a jwt in a cookie
// @Tags Auth
// @Accept  json
// @Produce
// @Param includeDeleted query bool false "Set to true to include soft deleted users" default(false)
// @Param full query bool false "Set to true to include full user details" default(false)
// @Success 200 {array} models.User "Successfully retrieved the list of users"
// @Failure 400 {object} map[string]interface{} "Error: No user records found"
// @Router /user/ [get]
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

		fmt.Printf("Token created: %s\n", tokenString)
		c.SetCookie("token", tokenString, 3600, "/", utils.GetEnvVariable("DOMAIN"), false, true)
		c.Redirect(http.StatusOK, "/")
	} else {
		c.String(http.StatusUnauthorized, "Invalid credentials")
	}
}
