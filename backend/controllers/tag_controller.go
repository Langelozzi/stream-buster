package controllers

import (
	"fmt"
	"strconv"

	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/models/search"

	"github.com/DOC-SWAP/Docswap-backend/services/interfaces"
	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagService interfaces.TagServiceInterface
}

func NewTagController(tagService interfaces.TagServiceInterface) *TagController {
	return &TagController{
		tagService: tagService,
	}
}

// GetAllTagsHandler retrieves all tags
// @Summary Retrieve all tags
// @Description Get all tags from the database
// @Tags tags
// @Produce  json
// @Param includeDeleted query bool false "Set to true to include soft deleted tags" default(false)
// @Param full query bool false "Set to true to include full tag details" default(false)
// @Success 200 {object} []models.Tag "Successfully retrieved the tag"
// @Failure 400 {object} map[string]interface{} "Error: Bad Request"
// @Router /tag/ [get]
func (contr *TagController) GetAllTagsHandler(c *gin.Context) {

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
	fullStr := c.DefaultQuery("full", "true")

	// convert the full query to a boolean
	full, err := strconv.ParseBool(fullStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid full query. Error: " + err.Error(),
		})
		return
	}

	// call the service
	tags, err := contr.tagService.GetAllTags(includeDeleted, full)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No tags found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, tags)
}

// GetTagHandler GetTag retrieves a tag record by its ID
// @Summary Retrieve tag record
// @Description get tag record by ID
// @Tags tags
// @Accept  json
// @Produce  json
// @Param id path int true "Tag ID"
// @Param includeDeleted query bool false "Set to true to include soft deleted tags" default(false)
// @Param full query bool false "Set to true to include full tag details" default(false)
// @Success 200 {object} models.Tag "Successfully retrieved the tag"
// @Failure 400 {object} map[string]interface{} "Error: Bad Request"
// @Router /tag/{id} [get]
func (contr *TagController) GetTagHandler(c *gin.Context) {
	// get the id from the request
	id, err := strconv.Atoi(c.Param("tagId"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Invalid id: %s. Error: %s", c.Query("id"), err.Error()),
		})
		return
	}

	// get the includeDeleted query
	includeDeletedStr := c.DefaultQuery("includeDeleted", "false")
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
	tag, err := contr.tagService.GetTag(id, includeDeleted, full)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Tag records not found" + err.Error(),
		})
		return
	}

	c.JSON(200, tag)
}

// CreateTagHandler creates a new tag
// @Summary Create a new tag record
// @Description Create a new tag record
// @Tags tags
// @Accept  json
// @Produce  json
// @Param tag body models.Tag true "Tag object"
// @Success 201 {object} map[string]interface{} "Successfully created the tag"
// @Failure 400 {object} map[string]interface{} "Error: Bad Request"
// @Router /tag/ [post]
func (contr *TagController) CreateTagHandler(c *gin.Context) {
	// create a new tag object
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid tag object " + err.Error(),
		})
		return
	}

	// call the service
	createdTag, err := contr.tagService.CreateTag(&tag)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create tag. Error: " + err.Error(),
		})
		return
	}

	c.JSON(201, createdTag)
}

// CreateTagsBulkHandler many new tag records
// @Summary Create multiple new tag records
// @Description Create multiple new tag records
// @Tags tags
// @Accept  json
// @Produce  json
// @Param tag body models.Tag true "Tag objects"
// @Success 201 {object} map[string]interface{} "Successfully created the tag records"
// @Failure 400 {object} map[string]interface{} "Error: Unable to create tag records"
// @Router /tag/ [post]
func (contr *TagController) CreateTagsBulkHandler(c *gin.Context) {
	// create a slice of tag objects
	var tags []models.Tag
	if err := c.ShouldBindJSON(&tags); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid tag object " + err.Error(),
		})
		return
	}

	// call the service
	createdTag, err := contr.tagService.CreateTagsBulk(tags)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create tag. Error: " + err.Error(),
		})
		return
	}

	c.JSON(201, createdTag)
}

// CreateTagsUserTagsBulkHandler many tags with user tags
// @Summary Create multiple new tag records
// @Description Create multiple new tag records
// @Tags tags
// @Accept  json
// @Produce  json
// @Param tag body models.Tag true "Tag objects"
// @Success 201 {object} map[string]interface{} "Successfully created the tag records"
// @Failure 400 {object} map[string]interface{} "Error: Unable to create tag records"
// @Router /tag/ [post]
func (contr *TagController) CreateTagsUserTagsBulkHandler(c *gin.Context) {
	// create a slice of tag objects
	var tags []models.Tag
	if err := c.ShouldBindJSON(&tags); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid tag object " + err.Error(),
		})
		return
	}
	// get the current user from the context
	currentUser, _ := c.Get("user")
	userId := currentUser.(*models.User).ID

	// call the service
	createdTag, err := contr.tagService.CreateTagsAndUserTagsBulk(tags, userId)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create tag. Error: " + err.Error(),
		})
		return
	}

	c.JSON(201, createdTag)
}

// DeleteTagHandler deletes an existing tag via its ID
// @Summary Delete a tag record
// @Description Delete a tag records via its ID
// @Tags tags
// @Accept  json
// @Produce  json
// @Param tagID path int true "Tag ID"
// @Param softDelete query bool true "Set to false to fully delete the tag record" default(true)
// @Success 204 {object} map[string]interface{} "Successfully deleted the tag record"
// @Failure 400 {object} map[string]interface{} "Error: Failed to delete the tag record"
// @Router /tag/{tagId} [delete]
func (contr *TagController) DeleteTagHandler(c *gin.Context) {
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
	softDelete, err := strconv.ParseBool(softDeleteStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid softDelete query. Error: " + err.Error(),
		})
		return
	}

	// call the service
	err = contr.tagService.DeleteTag(tagId, softDelete)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to delete tag record. Error: " + err.Error(),
		})
		return
	}

	c.JSON(204, gin.H{})
}

// DeleteTagsBulkHandler Deletes multiple tag records via tag ID
// @Summary Delete multiple tag records
// @Description Delete multiple tag records using tagID
// @Tags tags
// @Accept  json
// @Produce  json
// @Param tags body []models.Tag true "Tag objects that need to be deleted"
// @Param softDelete query bool true "Set to false to fully delete the tag records" default(true)
// @Success 204 {object} map[string]interface{} "Successfully deleted the tag records"
// @Failure 400 {object} map[string]interface{} "Error: Unable to delete tag records"
// @Router /tag/bulk [delete]
func (contr *TagController) DeleteTagsBulkHandler(c *gin.Context) {
	// get the tag id from the request
	var tags []models.Tag
	if err := c.ShouldBindJSON(&tags); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body. Error: " + err.Error(),
		})
		return
	}

	// get the softDelete query
	softDeleteStr := c.DefaultQuery("softDelete", "true")
	softDelete, err := strconv.ParseBool(softDeleteStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid softDelete query. Error: " + err.Error(),
		})
		return
	}

	// call the service
	err = contr.tagService.DeleteTagsBulk(tags, softDelete)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to delete tag records. Error: " + err.Error(),
		})
		return
	}

	c.JSON(204, gin.H{})
}

// SearchTagsHandler searches for tags with the given search parameters
// @Summary Search tags
// @Description Search for tags with the given search parameters
// @Tags tags
// @Accept  json
// @Produce  json
// @Param searchObj body search.Search true "Search object"
// @Success 200 {object} []models.Tag "Successfully retrieved the tag"
// @Failure 400 {object} map[string]interface{} "Error: Bad Request"
// @Router /tag/search [post]
func (contr *TagController) SearchTagsHandler(c *gin.Context) {
	// create a new search object
	var searchObj search.Search
	if err := c.ShouldBindJSON(&searchObj); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid search object " + err.Error(),
		})
		return
	}

	// call the service
	tags, err := contr.tagService.SearchTags(searchObj)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No tags found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, tags)
}
