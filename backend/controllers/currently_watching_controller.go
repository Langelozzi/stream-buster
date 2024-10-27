package controllers

import (
	"strconv"

	"github.com/STREAM-BUSTER/stream-buster/models/db"
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/gin-gonic/gin"
)

type CurrentlyWatchingController struct {
	service interfaces.CurrentlyWatchingServiceInterface
}

func NewCurrentlyWatchingController(service interfaces.CurrentlyWatchingServiceInterface) *CurrentlyWatchingController {
	return &CurrentlyWatchingController{
		service: service,
	}
}

// CreateCurrentlyWatchingHandler creates a new currently watching record
// @Summary Create a new currently watching record
// @Description create a new currently watching record
// @Tags currently-watching
// @Accept  json
// @Produce  json
// @Param watch body db.CurrentlyWatching true "CurrentlyWatching object that needs to be created"
// @Success 201 {object} db.CurrentlyWatching "Successfully created the currently watching record"
// @Failure 400 {object} map[string]interface{} "Error: Invalid request body"
// @Router /currently-watching/ [post]
func (contr *CurrentlyWatchingController) CreateCurrentlyWatchingHandler(c *gin.Context) {
	watch := &db.CurrentlyWatching{}
	if err := c.ShouldBindJSON(watch); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body. Error: " + err.Error(),
		})
		return
	}

	watch, err := contr.service.CreateCurrentlyWatching(watch)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create a currently watching record. Error: " + err.Error(),
		})
		return
	}

	c.JSON(201, watch)
}

// GetCurrentlyWatchingHandler retrieves a currently watching record
// @Summary Retrieve a currently watching record
// @Description get a currently watching record
// @Tags currently-watching
// @Accept  json
// @Produce  json
// @Param userID path int true "User ID"
// @Param mediaId path string true "Media ID"
// @Param includeDeleted query bool false "Set to false to exclude soft deleted record" default(false)
// @Success 200 {object} db.CurrentlyWatching "Successfully retrieved the currently watching record"
// @Failure 400 {object} map[string]interface{} "Error: Record not found"
// @Router /currently-watching/{userID}/{mediaId}/ [get]
func (contr *CurrentlyWatchingController) GetCurrentlyWatchingHandler(c *gin.Context) {
	idStr := c.Param("userID")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid user ID. Error: " + err.Error(),
		})
		return
	}

	mediaId := c.Param("mediaId")

	includeDeletedStr := c.DefaultQuery("includeDeleted", "false")
	includeDeleted, err := strconv.ParseBool(includeDeletedStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid includeDeleted query. Error: " + err.Error(),
		})
		return
	}

	watch, err := contr.service.GetCurrentlyWatchingById(uint(userID), mediaId, includeDeleted)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No currently watching records found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, watch)
}

// UpdateCurrentlyWatchingHandler updates a currently watching record
// @Summary Update a currently watching record
// @Description update a currently watching record
// @Tags currently-watching
// @Accept  json
// @Produce  json
// @Param watch body db.CurrentlyWatching true "CurrentlyWatching object that needs to be updated"
// @Success 200 {object} db.CurrentlyWatching "Successfully updated the currently watching record"
// @Failure 400 {object} map[string]interface{} "Error: Invalid request body"
// @Router /currently-watching/ [put]
func (contr *CurrentlyWatchingController) UpdateCurrentlyWatchingHandler(c *gin.Context) {
	watch := &db.CurrentlyWatching{}
	if err := c.ShouldBindJSON(watch); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body. Error: " + err.Error(),
		})
		return
	}

	updatedWatch, err := contr.service.UpdateCurrentlyWatching(watch)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to update currently watching record. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, updatedWatch)
}
