package controllers

import (
	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/services/interfaces"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserDocumentController struct {
	service interfaces.UserDocumentServiceInterface
}

func NewUserDocumentController(service interfaces.UserDocumentServiceInterface) *UserDocumentController {
	return &UserDocumentController{
		service: service,
	}
}

// GetAllUserDocumentsHandler Retrieves all user document records
// @Summary Retrieve all user document records
// @Description Get all user document records
// @Tags userdocuments
// @Accept  json
// @Produce  json
// @Param includeDeleted query bool false "Set to true to include soft deleted user documents" default(false)
// @Param full query bool false "Set to true to include full user and document details" default(false)
// @Success 200 {array} models.UserDocument "Successfully retrieved the list of user document records"
// @Failure 400 {object} map[string]interface{} "Error: No user document records found"
// @Router /userdocument/ [get]
func (contr *UserDocumentController) GetAllUserDocumentsHandler(c *gin.Context) {
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
	records, err := contr.service.GetAllUserDocuments(includeDeleted, full)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No user document records found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, records)
}

// GetUserDocumentHandler Retrieves a user document record by its user ID and document ID
// @Summary Retrieve a user document record by its ID
// @Description Get a user document record by its ID
// @Tags userdocuments
// @Accept  json
// @Produce  json
// @Param userId path int true "User ID"
// @Param documentId path int true "Document ID"
// @Param includeDeleted query bool false "Set to true to include soft deleted user documents" default(false)
// @Param full query bool false "Set to true to include full user and document details" default(false)
// @Success 200 {object} models.UserDocument "Successfully retrieved the user document record"
// @Failure 400 {object} map[string]interface{} "Error: No user document record found"
// @Router /userdocument/{userId}/{documentId} [get]
func (contr *UserDocumentController) GetUserDocumentHandler(c *gin.Context) {
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
	record, err := contr.service.GetUserDocument(userId, documentId, includeDeleted, full)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No user document records found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, record)
}

// CreateUserDocumentHandler Creates a new user document record
// @Summary Create a new user document record
// @Description Create a new user document record
// @Tags userdocuments
// @Accept  json
// @Produce  json
// @Param userDocument body models.UserDocument true "User Document object that needs to be created"
// @Success 201 {object} models.UserDocument "Successfully created the user document record"
// @Failure 400 {object} map[string]interface{} "Error: Unable to create user document record"
// @Router /userdocument/ [post]
func (contr *UserDocumentController) CreateUserDocumentHandler(c *gin.Context) {
	// get the body of the request
	userDocument := &models.UserDocument{}
	if err := c.ShouldBindJSON(userDocument); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body. Error: " + err.Error(),
		})
		return
	}

	// call the service
	userDocument, err := contr.service.CreateUserDocument(userDocument)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create new UserDocument record. Error: " + err.Error(),
		})
		return
	}

	c.JSON(201, userDocument)
}

// CreateUserDocumentsBulkHandler Creates multiple new user document records
// @Summary Create multiple new user document records
// @Description Create multiple new user document records
// @Tags userdocuments
// @Accept  json
// @Produce  json
// @Param userDocuments body []models.UserDocument true "User Document objects that need to be created"
// @Success 201 {object} []models.UserDocument "Successfully created the user document records"
// @Failure 400 {object} map[string]interface{} "Error: Unable to create user document records"
// @Router /userdocument/bulk [post]
func (contr *UserDocumentController) CreateUserDocumentsBulkHandler(c *gin.Context) {
	// get the body of the request
	var userDocuments []models.UserDocument
	if err := c.ShouldBindJSON(&userDocuments); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body. Error: " + err.Error(),
		})
		return
	}

	// call the service
	userDocuments, err := contr.service.CreateUserDocumentsBulk(userDocuments)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create new UserDocument records. Error: " + err.Error(),
		})
		return
	}

	c.JSON(201, userDocuments)
}

// DeleteUserDocumentHandler Deletes a user document record by its user ID and document ID
// @Summary Delete a user document record by its ID
// @Description Delete a user document record by its ID
// @Tags userdocuments
// @Accept  json
// @Produce  json
// @Param userId path int true "User ID"
// @Param documentId path int true "Document ID"
// @Param softDelete query bool true "Set to false to fully delete the user document record" default(true)
// @Success 204 {object} map[string]interface{} "Successfully deleted the user document record"
// @Failure 400 {object} map[string]interface{} "Error: Unable to delete user document record"
// @Router /userdocument/{userId}/{documentId} [delete]
func (contr *UserDocumentController) DeleteUserDocumentHandler(c *gin.Context) {
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
	err = contr.service.DeleteUserDocument(userId, documentId, softDelete)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to delete UserDocument record. Error: " + err.Error(),
		})
		return
	}

	c.JSON(204, gin.H{})
}

// DeleteUserDocumentsBulkHandler Deletes multiple user document records by their user ID and document ID
// @Summary Delete multiple user document records by their ID
// @Description Delete multiple user document records by their ID
// @Tags userdocuments
// @Accept  json
// @Produce  json
// @Param userDocuments body []models.UserDocument true "User Document objects that need to be deleted"
// @Param softDelete query bool true "Set to false to fully delete the user document records" default(true)
// @Success 204 {object} map[string]interface{} "Successfully deleted the user document records"
// @Failure 400 {object} map[string]interface{} "Error: Unable to delete user document records"
// @Router /userdocument/bulk [delete]
func (contr *UserDocumentController) DeleteUserDocumentsBulkHandler(c *gin.Context) {
	// get the body of the request
	var userDocuments []models.UserDocument
	if err := c.ShouldBindJSON(&userDocuments); err != nil {
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
	err = contr.service.DeleteUserDocumentsBulk(userDocuments, softDelete)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to delete UserDocument records. Error: " + err.Error(),
		})
		return
	}

	c.JSON(204, gin.H{})
}
