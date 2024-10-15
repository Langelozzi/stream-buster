package controllers

import (
	"strconv"

	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/services/interfaces"
	"github.com/gin-gonic/gin"
)

type UserTagController struct {
	service interfaces.UserTagServiceInterface
}

func NewUserTagController(service interfaces.UserTagServiceInterface) *UserTagController {
	return &UserTagController{
		service: service,
	}
}

// GetAllUserTagsHandler Retrieves all user Tag records
// @Summary Retrieve all user Tag records
// @Description Get all user Tag records
// @Tags UserTags
// @Accept  json
// @Produce  json
// @Param includeDeleted query bool false "Set to true to include soft deleted user Tags" default(false)
// @Param full query bool false "Set to true to include full user and Tag details" default(false)
// @Success 200 {array} models.UserTag "Successfully retrieved the list of user Tag records"
// @Failure 400 {object} map[string]interface{} "Error: No user Tag records found"
// @Router /usertag/ [get]
func (contr *UserTagController) GetAllUserTagsHandler(c *gin.Context) {
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

	// call the service
	records, err := contr.service.GetAllUserTags(includeDeleted, full)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No user Tag records found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, records)
}

// GetUserTagHandler Retrieves a user Tag record by its user ID and Tag ID
// @Summary Retrieve a user Tag record by its ID
// @Description Get a user Tag record by its ID
// @Tags UserTags
// @Accept  json
// @Produce  json
// @Param userId path int true "User ID"
// @Param TagId path int true "Tag ID"
// @Param includeDeleted query bool false "Set to true to include soft deleted user Tags" default(false)
// @Param full query bool false "Set to true to include full user and Tag details" default(false)
// @Success 200 {object} models.UserTag "Successfully retrieved the user Tag record"
// @Failure 400 {object} map[string]interface{} "Error: No user Tag record found"
// @Router /usertag/{userId}/{tagId} [get]
func (contr *UserTagController) GetUserTagHandler(c *gin.Context) {
	// get the user id from the request
	userIdStr := c.Param("userId")
	// convert the id to an integer
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid user ID. Error: " + err.Error(),
		})
		return
	}

	// get the Tag id from the request
	tagIdStr := c.Param("tagId")
	// convert the id to an integer
	tagId, err := strconv.Atoi(tagIdStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Tag ID. Error: " + err.Error(),
		})
		return
	}

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

	// call the service
	record, err := contr.service.GetUserTag(userId, tagId, includeDeleted, full)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No user Tag records found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, record)
}

// CreateUserTagHandler Creates a new user Tag record
// @Summary Create a new user Tag record
// @Description Create a new user Tag record
// @Tags UserTags
// @Accept  json
// @Produce  json
// @Param UserTag body models.UserTag true "User Tag object that needs to be created"
// @Success 201 {object} models.UserTag "Successfully created the user Tag record"
// @Failure 400 {object} map[string]interface{} "Error: Unable to create user Tag record"
// @Router /usertag/ [post]
func (contr *UserTagController) CreateUserTagHandler(c *gin.Context) {
	// get the body of the request
	userTag := &models.UserTag{}
	if err := c.ShouldBindJSON(userTag); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body. Error: " + err.Error(),
		})
		return
	}

	// call the service
	userTag, err := contr.service.CreateUserTag(userTag)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create new UserTag record. Error: " + err.Error(),
		})
		return
	}

	c.JSON(201, userTag)
}

// CreateUserTagsBulkHandler Creates multiple new user Tag records
// @Summary Create multiple new user Tag records
// @Description Create multiple new user Tag records
// @Tags UserTags
// @Accept  json
// @Produce  json
// @Param UserTags body []models.UserTag true "User Tag objects that need to be created"
// @Success 201 {object} []models.UserTag "Successfully created the user Tag records"
// @Failure 400 {object} map[string]interface{} "Error: Unable to create user Tag records"
// @Router /usertag/bulk [post]
func (contr *UserTagController) CreateUserTagsBulkHandler(c *gin.Context) {
	// get the body of the request
	var userTags []models.UserTag
	if err := c.ShouldBindJSON(&userTags); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body. Error: " + err.Error(),
		})
		return
	}

	// call the service
	userTags, err := contr.service.CreateUserTagsBulk(userTags)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create new UserTag records. Error: " + err.Error(),
		})
		return
	}

	c.JSON(201, userTags)
}

// DeleteUserTagHandler Deletes a user Tag record by its user ID and Tag ID
// @Summary Delete a user Tag record by its ID
// @Description Delete a user Tag record by its ID
// @Tags UserTags
// @Accept  json
// @Produce  json
// @Param userId path int true "User ID"
// @Param TagId path int true "Tag ID"
// @Param softDelete query bool true "Set to false to fully delete the user Tag record" default(true)
// @Success 204 {object} map[string]interface{} "Successfully deleted the user Tag record"
// @Failure 400 {object} map[string]interface{} "Error: Unable to delete user Tag record"
// @Router /usertag/{userId}/{TagId} [delete]
func (contr *UserTagController) DeleteUserTagHandler(c *gin.Context) {
	// get the user id from the request
	userIdStr := c.Param("userId")
	// convert the id to an integer
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid user ID. Error: " + err.Error(),
		})
		return
	}

	// get the tag id from the request
	tagIdStr := c.Param("tagId")
	// convert the id to an integer
	tagId, err := strconv.Atoi(tagIdStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid tag ID. Error: " + err.Error(),
		})
		return
	}

	// get the softDelete query
	softDeleteStr := c.DefaultQuery("softDelete", "true")
	// convert the softDelete query to a boolean
	softDelete, err := strconv.ParseBool(softDeleteStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid softDelete query. Error: " + err.Error(),
		})
		return
	}

	// call the service
	err = contr.service.DeleteUserTag(userId, tagId, softDelete)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to delete UserTag record. Error: " + err.Error(),
		})
		return
	}

	c.JSON(204, gin.H{})
}

// DeleteUserTagsBulkHandler Deletes multiple user Tag records by their user ID and Tag ID
// @Summary Delete multiple user Tag records by their ID
// @Description Delete multiple user Tag records by their ID
// @Tags UserTags
// @Accept  json
// @Produce  json
// @Param UserTags body []models.UserTag true "User Tag objects that need to be deleted"
// @Param softDelete query bool true "Set to false to fully delete the user Tag records" default(true)
// @Success 204 {object} map[string]interface{} "Successfully deleted the user Tag records"
// @Failure 400 {object} map[string]interface{} "Error: Unable to delete user Tag records"
// @Router /usertag/bulk [delete]
func (contr *UserTagController) DeleteUserTagsBulkHandler(c *gin.Context) {
	// get the body of the request
	var userTags []models.UserTag
	if err := c.ShouldBindJSON(&userTags); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body. Error: " + err.Error(),
		})
		return
	}

	// get the softDelete query
	softDeleteStr := c.DefaultQuery("softDelete", "true")
	// convert the softDelete query to a boolean
	softDelete, err := strconv.ParseBool(softDeleteStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid softDelete query. Error: " + err.Error(),
		})
		return
	}

	// call the service
	err = contr.service.DeleteUserTagsBulk(userTags, softDelete)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to delete UserTag records. Error: " + err.Error(),
		})
		return
	}

	c.JSON(204, gin.H{})
}
