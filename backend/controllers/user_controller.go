package controllers

import (
	"strconv"

	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/utils/auth"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service interfaces.UserServiceInterface
}

func NewUserController(service interfaces.UserServiceInterface) *UserController {
	return &UserController{
		service: service,
	}
}

// GetAllUsersHandler retrieves all users
// @Summary Retrieve all users
// @Description get all user records
// @Tags users
// @Accept  json
// @Produce  json
// @Param includeDeleted query bool false "Set to true to include soft deleted users" default(false)
// @Param full query bool false "Set to true to include full user details" default(false)
// @Success 200 {array} models.User "Successfully retrieved the list of users"
// @Failure 400 {object} map[string]interface{} "Error: No user records found"
// @Router /user/ [get]
func (contr *UserController) GetAllUsersHandler(c *gin.Context) {
	// get the includeDeleted query
	includeDeletedStr := c.DefaultQuery("includeDeleted", "false")
	// convert the includeDeleted query to a boolean
	includeDeleted, err := strconv.ParseBool(includeDeletedStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid includeDeleted query. Error: " + err.Error(),
		})
		return
	}

	// get the full query
	fullStr := c.DefaultQuery("full", "false")
	// convert the full query to a boolean
	full, err := strconv.ParseBool(fullStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid full query. Error: " + err.Error(),
		})
		return
	}

	// call the Service
	users, err := contr.service.GetAllUsers(includeDeleted, full)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No user records found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

// GetUserHandler retrieves a user
// @Summary Retrieve a user
// @Description get a user record
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param includeDeleted query bool false "Set to true to include soft deleted user" default(false)
// @Param full query bool false "Set to true to include full user details" default(false)
// @Success 200 {object} models.User "Successfully retrieved the user"
// @Failure 400 {object} map[string]interface{} "Error: No user records found"
// @Router /user/{id}/ [get]
func (contr *UserController) GetUserHandler(c *gin.Context) {
	// get the id from the request
	idStr := c.Param("id")
	// convert the id to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid user ID. Error: " + err.Error(),
		})
		return
	}

	//get the includeDeleted query
	includeDeletedStr := c.DefaultQuery("includeDeleted", "false")
	// convert the includeDeleted query to a boolean
	includeDeleted, err := strconv.ParseBool(includeDeletedStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid includeDeleted query. Error: " + err.Error(),
		})
		return
	}

	// get the full query
	fullStr := c.DefaultQuery("full", "false")
	// convert the full query to a boolean
	full, err := strconv.ParseBool(fullStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid full query. Error: " + err.Error(),
		})
		return
	}

	// call the Service
	user, err := contr.service.GetUser(id, includeDeleted, full)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No user records found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

// GetCurrentUserHandler retrieves the current user
// @Summary Retrieve the current user
// @Description get the current user record
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User "Successfully retrieved the current user"
// @Failure 400 {object} map[string]interface{} "Error: No user records found"
// @Router /user/current/ [get]
func (contr *UserController) GetCurrentUserHandler(c *gin.Context) {
	// Get the current user from the context
	var ctxUser models.User
	ctxUser, err := auth.GetUserFromContext(c)
	userID := ctxUser.ID

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error Validating user: Error Cannot Get Id from Context",
		})
		return
	}

	//get the includeDeleted query
	includeDeletedStr := c.DefaultQuery("includeDeleted", "false")
	// convert the includeDeleted query to a boolean
	includeDeleted, err := strconv.ParseBool(includeDeletedStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid includeDeleted query. Error: " + err.Error(),
		})
		return
	}

	// get the full query
	fullStr := c.DefaultQuery("full", "false")
	// convert the full query to a boolean
	full, err := strconv.ParseBool(fullStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid full query. Error: " + err.Error(),
		})
		return
	}

	user, err := contr.service.GetUser(int(userID), includeDeleted, full)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting full user record",
		})
		return
	}

	c.JSON(200, user)
}

// CreateUserHandler creates a new user
// @Summary Create a new user
// @Description create a new user record
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User object that needs to be created"
// @Success 201 {object} models.User "Successfully created the user"
// @Failure 400 {object} map[string]interface{} "Error: Invalid request body"
// @Router /user/ [post]
func (contr *UserController) CreateUserHandler(c *gin.Context) {
	// get the body of the request
	user := &models.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body. Error: " + err.Error(),
		})
		return
	}

	// call the Service
	user, err := contr.service.CreateUser(user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No user records found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(201, user)
}

// UpdateUserHandler updates a user
// @Summary Update a user
// @Description update a user record
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body models.User true "User object that needs to be updated"
// @Success 200 {object} models.User "Successfully updated the user"
// @Failure 400 {object} map[string]interface{} "Error: Invalid request body"
// @Router /user/{id}/ [put]
func (contr *UserController) UpdateUserHandler(c *gin.Context) {
	// get the id from the request
	idStr := c.Param("id")
	// convert the id to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid user ID. Error: " + err.Error(),
		})
		return
	}

	// get the body of the request
	user := &models.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body. Error: " + err.Error(),
		})
		return
	}

	// set the id of the user
	user.ID = uint64(id)

	// call the Service
	user, err = contr.service.UpdateUser(user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No user records found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

// UpdateCurrentUserHandler updates the profile of the currently authenticated user using external ID
// @Summary Update current user
// @Description update the authenticated user's record using external ID
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User "Successfully updated the current user"
// @Failure 400 {object} map[string]interface{} "Error: Invalid request body"
// @Failure 404 {object} map[string]interface{} "Error: User not found"
// @Router /user/current [put]
//func (contr *UserController) UpdateCurrentUserHandler(c *gin.Context) {
//	userInterface, exists := c.Get("user")
//	if !exists {
//		c.JSON(404, gin.H{"error": "User not found"})
//		return
//	}
//
//	user, ok := userInterface.(*models.User)
//	if !ok || user.ExternalUserID == "" {
//		c.JSON(404, gin.H{"error": "User not found"})
//		return
//	}
//
//	var updateData models.User
//	if err := c.ShouldBindJSON(&updateData); err != nil {
//		c.JSON(400, gin.H{"error": "Invalid request body"})
//		return
//	}
//
//	updatedUser, err := contr.Service.UpdateUserByExternalID(user.ExternalUserID, &updateData)
//	if err != nil {
//		c.JSON(500, gin.H{"error": "Failed to update user: " + err.Error()})
//		return
//	}
//
//	c.JSON(200, updatedUser)
//}

// DeleteUserHandler deletes a user
// @Summary Delete a user
// @Description delete a user record
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param softDelete query bool false "Set to false to disable soft delete" default(true)
// @Success 200 {object} map[string]interface{} "Successfully deleted the user"
// @Failure 400 {object} map[string]interface{} "Error: Invalid user ID"
// @Router /user/{id}/ [delete]
func (contr *UserController) DeleteUserHandler(c *gin.Context) {
	// get the id from the request
	idStr := c.Param("id")
	// convert the id to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid user ID. Error: " + err.Error(),
		})
		return
	}

	// get the softDelete query
	softDeleteStr := c.DefaultQuery("softDelete", "true")
	// convert the softDelete query to a boolean
	softDelete, err := strconv.ParseBool(softDeleteStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid useSoftDelete query. Error: " + err.Error(),
		})
		return
	}

	// call the Service
	err = contr.service.DeleteUser(id, softDelete)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No user records found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(204, gin.H{
		"message": "Successfully deleted the user",
	})
}

// GetUserUsageHandler retrieves the current user
// @Summary Retrieve the current user
// @Description get the current user record
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User "Successfully retrieved the current user"
// @Failure 400 {object} map[string]interface{} "Error: No user records found"
// @Router /user/current/ [get]
func (contr *UserController) GetUserUsageHandler(c *gin.Context) {
	// get the user id from the request
	idStr := c.Param("id")
	// convert the id to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid user ID. Error: " + err.Error(),
		})
		return
	}

	usage, err := contr.service.GetUserUsageByUserId(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting full usage record",
		})
		return
	}

	c.JSON(200, usage)
}
