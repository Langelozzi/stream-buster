package controllers

import (
	"net/http"
	"strconv"

	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service     interfaces.AuthServiceInterface
	userService interfaces.UserServiceInterface
}

func NewAuthController(service interfaces.AuthServiceInterface, userService interfaces.UserServiceInterface) *AuthController {
	return &AuthController{
		Service:     service,
		userService: userService,
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
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := contr.userService.GetUserByEmail(email, false, false)

	if err != nil {
		c.String(400, "User does not not exist")
	}

	validCredentials := contr.Service.CheckCredentials(email, password, user)

	if validCredentials {

		tokenString, err := contr.Service.CreateToken(email)

		if err != nil || tokenString == "" {
			c.String(http.StatusInternalServerError, "Error creating token")
			return
		}

		refreshTokenString, err := contr.Service.CreateRefreshToken(email)

		if err != nil || refreshTokenString == "" {
			c.String(http.StatusInternalServerError, "Error creating refreshToken")
			return
		}

		maxRefreshTokenAge, err := strconv.Atoi(utils.GetEnvVariable("REFRESH_TOKEN_EXPIRATION_TIME"))

		if err != nil {
			c.String(http.StatusInternalServerError, "Error fetching refresh token age")
		}

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

		c.String(http.StatusOK, "Authorized")

	} else {
		c.String(http.StatusUnauthorized, "Invalid Credentials")
	}
}

func (contr *AuthController) CreateUser(c *gin.Context) {
	email := c.PostForm("Email")
	password := c.PostForm("Password")
	firstName := c.PostForm("FirstName")
	lastName := c.PostForm("LastName")

	// Create the user object
	newUser := models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}

	createdUser, err := contr.userService.CreateUser(&newUser)
	if err != nil {
		c.String(400, "Error Creating user")
	}

	c.JSON(201, createdUser)
}

func (contr *AuthController) TestAuthMiddleware(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "No user found"})
	}
	c.JSON(200, user)
}
