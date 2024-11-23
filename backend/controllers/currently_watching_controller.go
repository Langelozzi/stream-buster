package controllers

import (
	"net/http"
	"strconv"

	"github.com/STREAM-BUSTER/stream-buster/models/db"
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/utils/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5/pgconn"
)

type CurrentlyWatchingController struct {
	service interfaces.CurrentlyWatchingServiceInterface
}

func NewCurrentlyWatchingController(service interfaces.CurrentlyWatchingServiceInterface) *CurrentlyWatchingController {
	return &CurrentlyWatchingController{
		service: service,
	}
}

func (contr *CurrentlyWatchingController) Test(c *gin.Context) {
	c.String(http.StatusOK, "success")
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

	user, err := auth.GetUserFromContext(c)

	if uint64(user.ID) != uint64(watch.UserID) {
		c.JSON(401, gin.H{
			"message": "Error: cannot verify user",
		})
		return
	}

	watch, err = contr.service.CreateCurrentlyWatching(watch)
	if pgError, ok := err.(*pgconn.PgError); ok {
		if pgError.Code == "23505" {
			c.JSON(400, gin.H{
				"duplicateKey": true,
				"message":      "Cannot Create Record; already exists",
			})
			return
		} else {
			c.JSON(400, gin.H{
				"message": "Failed to create a currently watching record. PostgreSQL Error Code: " + pgError.Code,
			})
			return
		}

	} else if err != nil {

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

	includeDeletedStr := c.DefaultQuery("includeDeleted", "false")
	includeDeleted, err := strconv.ParseBool(includeDeletedStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid includeDeleted query. Error: " + err.Error(),
		})
		return
	}

	watch, err := contr.service.GetCurrentlyWatchingByUserId(uint(userID), includeDeleted)
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

func (contr *CurrentlyWatchingController) GetAllCurrentlyWatchingHandler(c *gin.Context) {
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{
			"message": "Error: cannot verify user",
		})
		return
	}

	id, ok := claims.(jwt.MapClaims)["id"].(float64)
	if !ok {
		c.JSON(401, gin.H{
			"message": "Error: cannot verify user",
		})
		return
	}

	// Parse includeDeleted query parameter
	includeDeletedStr := c.DefaultQuery("includeDeleted", "false")
	includeDeleted, err := strconv.ParseBool(includeDeletedStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid includeDeleted query. Error: " + err.Error(),
		})
		return
	}

	// Retrieve all currently watching records for the authenticated user
	watches, err := contr.service.GetCurrentlyWatchingByUserId(uint(id), includeDeleted)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No currently watching records found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, watches)
}

func (contr *CurrentlyWatchingController) GetWatchlist(c *gin.Context) {
	user, err := auth.GetUserFromContext(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error getting user information. Error: " + err.Error(),
		})
		return
	}
	// Retrieve all currently watching records for the authenticated user
	watches, err := contr.service.GetWatchlist(uint(user.ID))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No currently watching records found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, watches)
}

func (contr *CurrentlyWatchingController) DeleteCurrentlyWatchingHandler(c *gin.Context) {
	user, err := auth.GetUserFromContext(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error getting user information. Error: " + err.Error(),
		})
		return
	}

	mediaIdUint, err := strconv.ParseUint(c.Param("mediaId"), 10, 64)

	err = contr.service.DeleteCurrentlyWatching(uint(user.ID), uint(mediaIdUint))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error deleting Currrently watching" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Currently watching successfully deleted",
	})

}
