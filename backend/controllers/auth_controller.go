package controllers

import (
	"net/http"
	"strconv"
	"strings"

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
// @Tags auth
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {string} string "Successfully logged in, JWT set in cookie"
// @Failure 400 {object} map[string]interface{} "Invalid username or password"
// @Router /auth/login [post]
func (contr *AuthController) LoginUser(c *gin.Context) {
	email := strings.ToLower(c.PostForm("email"))
	password := c.PostForm("password")

	user, err := contr.userService.GetUserByEmail(email, false, true)

	if err != nil || user == nil {
		c.String(400, "User does not not exist")
	}

	validCredentials := contr.Service.CheckCredentials(password, user)

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

		c.JSON(200, gin.H{
			"user":  user,
			"token": tokenString,
		})

	} else {
		c.String(http.StatusUnauthorized, "Invalid Credentials")
	}
}

// RegisterUser registers a new user
// @Summary Register a new user
// @Description Register a user with email, password, first name, and last name
// @Tags auth
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param email formData string true "Email of the user"
// @Param password formData string true "Password of the user"
// @Param firstName formData string true "First name of the user"
// @Param lastName formData string true "Last name of the user"
// @Success 201 {object} models.User "Successfully created the user"
// @Failure 400 {string} string "Error: User creation failed"
// @Router /auth/register [post]
func (contr *AuthController) RegisterUser(c *gin.Context) {
	email := strings.ToLower(c.PostForm("email"))
	password := c.PostForm("password")
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")

	// Create the user object
	newUser := models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}

	createdUser, err := contr.Service.Register(newUser)
	if err != nil {
		c.String(400, "Error Creating user")
	}

	c.JSON(201, createdUser)
}

// LogoutUser logs out the user by clearing authentication tokens
// @Summary Log out user
// @Description Clear all authentication tokens from cookies and log out the user
// @Tags auth
// @Produce  json
// @Success 200 {object} map[string]string "Logged out successfully"
// @Router /auth/logout [post]
func (contr *AuthController) LogoutUser(c *gin.Context) {
	// Clear the refresh token cookie
	c.SetCookie(
		"refreshToken",                 // Name of the cookie
		"",                             // Clear the value
		-1,                             // MaxAge = -1 means the cookie expires immediately
		"/",                            // Path
		utils.GetEnvVariable("DOMAIN"), // Domain
		false,                          // Secure flag
		false,                          // HttpOnly flag
	)

	// Clear the access token cookie if you are using one
	c.SetCookie(
		"token",                        // Name of the cookie
		"",                             // Clear the value
		-1,                             // MaxAge = -1 means the cookie expires immediately
		"/",                            // Path
		utils.GetEnvVariable("DOMAIN"), // Domain
		false,                          // Secure flag
		false,                          // HttpOnly flag
	)

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (contr *AuthController) TestAuthMiddleware(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "No user found"})
	}
	c.JSON(200, user)
}
