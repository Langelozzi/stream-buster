package controllers

import (
	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/services/interfaces"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DocumentTagController struct {
	service interfaces.DocumentTagServiceInterface
}

func NewDocumentTagController(service interfaces.DocumentTagServiceInterface) *DocumentTagController {
	return &DocumentTagController{
		service: service,
	}
}

// GetAllDocumentTagsHandler Retrieves all document tag records
// @Summary Retrieve all document tag records
// @Description Get all document tag records
// @Tags documenttags
// @Accept  json
// @Produce  json
// @Param includeDeleted query bool false "Set to true to include soft deleted document tags" default(false)
// @Param full query bool false "Set to true to include full tag and document details" default(false)
// @Success 200 {array} models.DocumentTag "Successfully retrieved the list of document tag records"
// @Failure 400 {object} map[string]interface{} "Error: No document tag records found"
// @Router /documenttag/ [get]
func (contr *DocumentTagController) GetAllDocumentTagsHandler(c *gin.Context) {
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
	records, err := contr.service.GetAllDocumentTags(includeDeleted, full)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No document tag records found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, records)
}

// GetDocumentTagHandler Retrieves a document tag record by its document ID and tag ID
// @Summary Retrieve a document tag record by its ID
// @Description Get a document tag record by its ID
// @Tags documenttags
// @Accept  json
// @Produce  json
// @Param documentId path int true "Document ID"
// @Param tagId path int true "Tag ID"
// @Param includeDeleted query bool false "Set to true to include soft deleted document tags" default(false)
// @Param full query bool false "Set to true to include full document and tag details" default(false)
// @Success 200 {object} models.DocumentTag "Successfully retrieved the document tag record"
// @Failure 400 {object} map[string]interface{} "Error: No document tag record found"
// @Router /documenttag/{documentId}/{tagId} [get]
func (contr *DocumentTagController) GetDocumentTagHandler(c *gin.Context) {
	// get the document id from the request
	documentIdStr := c.Param("documentId")
	// convert the id to an integer
	documentId, err := strconv.Atoi(documentIdStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid document ID. Error: " + err.Error(),
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
	record, err := contr.service.GetDocumentTag(documentId, tagId, includeDeleted, full)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No document tag records found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, record)
}

// CreateDocumentTagHandler Creates a new document tag record
// @Summary Create a new document tag record
// @Description Create a new document tag record
// @Tags documenttags
// @Accept  json
// @Produce  json
// @Param documentTag body models.DocumentTag true "Document Tag object that needs to be created"
// @Success 201 {object} models.DocumentTag "Successfully created the document tag record"
// @Failure 400 {object} map[string]interface{} "Error: Unable to create document tag record"
// @Router /documenttag/ [post]
func (contr *DocumentTagController) CreateDocumentTagHandler(c *gin.Context) {
	// get the body of the request
	documentTag := &models.DocumentTag{}
	if err := c.ShouldBindJSON(documentTag); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body. Error: " + err.Error(),
		})
		return
	}

	// call the service
	documentTag, err := contr.service.CreateDocumentTag(documentTag)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create new DocumentTag record. Error: " + err.Error(),
		})
		return
	}

	c.JSON(201, documentTag)
}

// CreateDocumentTagsBulkHandler Creates multiple new document tag records
// @Summary Create multiple new document tag records
// @Description Create multiple new document tag records
// @Tags documenttags
// @Accept  json
// @Produce  json
// @Param documentTags body []models.DocumentTag true "Document Tag objects that need to be created"
// @Success 201 {object} []models.DocumentTag "Successfully created the document tag records"
// @Failure 400 {object} map[string]interface{} "Error: Unable to create document tag records"
// @Router /documenttag/bulk [post]
func (contr *DocumentTagController) CreateDocumentTagsBulkHandler(c *gin.Context) {
	// get the body of the request
	var documentTags []models.DocumentTag
	if err := c.ShouldBindJSON(&documentTags); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body. Error: " + err.Error(),
		})
		return
	}

	// call the service
	documentTags, err := contr.service.CreateDocumentTagsBulk(documentTags)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create new DocumentTag records. Error: " + err.Error(),
		})
		return
	}

	c.JSON(201, documentTags)
}

// DeleteDocumentTagHandler Deletes a document tag record by its document ID and tag ID
// @Summary Delete a document tag record by its ID
// @Description Delete a document tag record by its ID
// @Tags documenttags
// @Accept  json
// @Produce  json
// @Param documentId path int true "Document ID"
// @Param tagId path int true "Tag ID"
// @Param softDelete query bool true "Set to false to fully delete the document tag record" default(true)
// @Success 204 {object} map[string]interface{} "Successfully deleted the document tag record"
// @Failure 400 {object} map[string]interface{} "Error: Unable to delete document tag record"
// @Router /documenttag/{documentId}/{tagId} [delete]
func (contr *DocumentTagController) DeleteDocumentTagHandler(c *gin.Context) {
	// get the document id from the request
	documentIdStr := c.Param("documentId")
	// convert the id to an integer
	documentId, err := strconv.Atoi(documentIdStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid document ID. Error: " + err.Error(),
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
	err = contr.service.DeleteDocumentTag(documentId, tagId, softDelete)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to delete DocumentTag record. Error: " + err.Error(),
		})
		return
	}

	c.JSON(204, gin.H{})
}

// DeleteDocumentTagsBulkHandler Deletes multiple document tag records by their document ID and tagId
// @Summary Delete multiple document tag records by their ID
// @Description Delete multiple document tag records by their ID
// @Tags documenttags
// @Accept  json
// @Produce  json
// @Param documentTags body []models.DocumentTag true "Document Tag objects that need to be deleted"
// @Param softDelete query bool true "Set to false to fully delete the document tag records" default(true)
// @Success 204 {object} map[string]interface{} "Successfully deleted the document tag records"
// @Failure 400 {object} map[string]interface{} "Error: Unable to delete document tag records"
// @Router /documenttag/bulk [delete]
func (contr *DocumentTagController) DeleteDocumentTagsBulkHandler(c *gin.Context) {
	// get the body of the request
	var documentTags []models.DocumentTag
	if err := c.ShouldBindJSON(&documentTags); err != nil {
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
	err = contr.service.DeleteDocumentTagsBulk(documentTags, softDelete)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to delete DocumentTag records. Error: " + err.Error(),
		})
		return
	}

	c.JSON(204, gin.H{})
}
