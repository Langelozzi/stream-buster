package controllers

import (
	"bytes"
	"fmt"
	"github.com/DOC-SWAP/Docswap-backend/models/search"
	"github.com/DOC-SWAP/Docswap-backend/utils"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/DOC-SWAP/Docswap-backend/models"

	"github.com/DOC-SWAP/Docswap-backend/services/interfaces"
	"github.com/gin-gonic/gin"
)

type DocumentController struct {
	documentService interfaces.DocumentServiceInterface
	documentTag     interfaces.DocumentTagServiceInterface
}

func NewDocumentController(documentService interfaces.DocumentServiceInterface, documentTagService interfaces.DocumentTagServiceInterface) *DocumentController {
	return &DocumentController{
		documentService: documentService,
		documentTag:     documentTagService,
	}
}

// GetAllDocumentsHandler retrieves all documents
// @Summary Retrieve all documents
// @Description Get all documents from the database
// @Tags documents
// @Produce  json
// @Param includeDeleted query bool false "Set to true to include soft deleted documents" default(false)
// @Param full query bool false "Set to true to include full tag details" default(false)
// @Success 200 {object} []models.Document "Successfully retrieved the document"
// @Failure 400 {object} map[string]interface{} "Error: Bad Request"
// @Router /document/ [get]
func (contr *DocumentController) GetAllDocumentsHandler(c *gin.Context) {

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
	documents, err := contr.documentService.GetAllDocuments(includeDeleted, full)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No documents found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, documents)
}

// GetDocumentHandler GetDocument retrieves a document record by its ID
// @Summary Retrieve document record
// @Description get document record by ID
// @Tags documents
// @Accept  json
// @Produce  json
// @Param id path int true "Document ID"
// @Param includeDeleted query bool false "Set to true to include soft deleted documents" default(false)
// @Param full query bool false "Set to true to include full tag details" default(false)
// @Success 200 {object} models.Document "Successfully retrieved the document"
// @Failure 400 {object} map[string]interface{} "Error: Bad Request"
// @Router /document/{id} [get]
func (contr *DocumentController) GetDocumentHandler(c *gin.Context) {
	// get the id from the request
	id, err := strconv.Atoi(c.Param("documentId"))
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
	document, err := contr.documentService.GetDocument(id, includeDeleted, full)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Document records not found" + err.Error(),
		})
		return
	}

	c.JSON(200, document)
}

// CreateDocumentHandler creates a new document
// @Summary Create a new document record
// @Description Create a new document record
// @Tags documents
// @Accept  json
// @Produce  json
// @Param document body models.Document true "Document object"
// @Success 201 {object} map[string]interface{} "Successfully created the document"
// @Failure 400 {object} map[string]interface{} "Error: Bad Request"
// @Router /document/ [post]
func (contr *DocumentController) CreateDocumentHandler(c *gin.Context) {
	// create a new document object
	var document models.Document
	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid document object " + err.Error(),
		})
		return
	}

	// call the service
	createdDocument, err := contr.documentService.CreateDocument(&document)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create document. Error: " + err.Error(),
		})
		return
	}

	c.JSON(201, createdDocument)
}

// CreateDocumentsBulkHandler many new document records
// @Summary Create multiple new document records
// @Description Create multiple new document records
// @Tags documents
// @Accept  json
// @Produce  json
// @Param document body models.Document true "Document objects"
// @Success 201 {object} map[string]interface{} "Successfully created the document records"
// @Failure 400 {object} map[string]interface{} "Error: Unable to create document records"
// @Router /document/ [post]
func (contr *DocumentController) CreateDocumentsBulkHandler(c *gin.Context) {
	// create a slice of document objects
	var documents []models.Document
	if err := c.ShouldBindJSON(&documents); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid document object " + err.Error(),
		})
		return
	}

	// call the service
	createdDocument, err := contr.documentService.CreateDocumentsBulk(documents)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create document. Error: " + err.Error(),
		})
		return
	}

	c.JSON(201, createdDocument)
}

// DeleteDocumentHandler deletes an existing document via its ID
// @Summary Delete a document record
// @Description Delete a document records via its ID
// @Tags documents
// @Accept  json
// @Produce  json
// @Param documentID path int true "Document ID"
// @Param softDelete query bool true "Set too false to fully delete the document record" default(true)
// @Success 204 {object} map[string]interface{} "Successfully deleted the document record"
// @Failure 400 {object} map[string]interface{} "Error: Failed to delete the document record"
// @Router /document/{documentId} [delete]
func (contr *DocumentController) DeleteDocumentHandler(c *gin.Context) {
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
	softDelete, err := strconv.ParseBool(softDeleteStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid softDelete query. Error: " + err.Error(),
		})
		return
	}

	// call the service
	err = contr.documentService.DeleteDocument(documentId, softDelete)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to delete document record. Error: " + err.Error(),
		})
		return
	}

	c.JSON(204, gin.H{})
}

// DeleteDocumentsBulkHandler Deletes multiple document records via document ID
// @Summary Delete multiple document records
// @Description Delete multiple document records using documentID
// @Tags documents
// @Accept  json
// @Produce  json
// @Param documents body []models.Document true "Document objects that need to be deleted"
// @Param softDelete query bool true "Set too false to fully delete the document records" default(true)
// @Success 204 {object} map[string]interface{} "Successfully deleted the document records"
// @Failure 400 {object} map[string]interface{} "Error: Unable to delete document records"
// @Router /document/bulk [delete]
func (contr *DocumentController) DeleteDocumentsBulkHandler(c *gin.Context) {
	// get the document id from the request
	var documents []models.Document
	if err := c.ShouldBindJSON(&documents); err != nil {
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
	err = contr.documentService.DeleteDocumentsBulk(documents, softDelete)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to delete document records. Error: " + err.Error(),
		})
		return
	}

	c.JSON(204, gin.H{})
}

// WriteDocument uploads a document
// @Summary Upload document
// @Description Upload a new document
// @Tags documents
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "Document file"
// @Success 200 {object} map[string]interface{} "Successfully uploaded the document"
// @Failure 400 {object} map[string]interface{} "Error: Bad Request"
// @Router /document/upload [post]
//
//	func (contr *DocumentController) WriteDocument(c *gin.Context) {
//		// get the document file from the request
//		form, _ := c.MultipartForm()
//		files := form.File["files[]"]
//
//		// get the current user from the context
//		currentUser, _ := c.Get("user")
//		userId := currentUser.(*models.User).ID
//
//		// create an array of document
//		var documents []models.Document
//		var successfulUploads int = 0
//
//		for _, file := range files {
//			openedFile, err := file.Open()
//			if err != nil {
//				log.Println(err)
//				c.JSON(400, gin.H{
//					"message": "Failed to open the file",
//				})
//				return
//			}
//			defer func(openedFile multipart.File) {
//				err := openedFile.Close()
//				if err != nil {
//
//				}
//			}(openedFile)
//
//			// call the service
//			docObj := models.Document{
//				FileName: file.Filename,
//				FilePath: utils.GetEnvVariable("BLOB_CONTAINER"),
//			}
//			document, err := contr.documentService.UploadDocument(userId, &docObj, openedFile)
//			if err != nil {
//				document = &models.Document{
//					FileName: file.Filename,
//					FilePath: "ERROR UPLOADING FILE",
//				}
//				documents = append(documents, *document)
//				continue
//			}
//			documents = append(documents, *document)
//			successfulUploads++
//		}
//		if successfulUploads != len(files) {
//			c.JSON(400, gin.H{
//				"message":   "Failed to upload the files",
//				"documents": documents,
//			})
//			return
//
//		}
//
//		// Return the response
//		c.JSON(200, gin.H{
//			"message":   "Some files uploaded successfully.",
//			"documents": documents,
//		})
//	}
//
// WriteDocument uploads a document
// @Summary Upload document
// @Description Upload a new document
// @Tags documents
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "Document file"
// @Success 200 {object} map[string]interface{} "Successfully uploaded the document"
// @Failure 400 {object} map[string]interface{} "Error: Bad Request"
// @Router /document/upload [post]

func (contr *DocumentController) WriteSingleDocument(c *gin.Context) {
	// get the document file from the request
	form, _ := c.MultipartForm()
	files := form.File["file"]

	// get the current user from the context
	currentUser, _ := c.Get("user")
	userId := currentUser.(*models.User).ID

	// open the file
	file := files[0]
	openedFile, err := file.Open()
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Failed to open the file",
		})
		return
	}
	// close the file
	defer func(openedFile multipart.File) {
		err := openedFile.Close()
		if err != nil {
			c.JSON(400, gin.H{"message": "Failed to close the file"})

		}
	}(openedFile)

	// get the category ID
	uInt64, err := strconv.ParseUint(form.Value["category_id"][0], 10, 64)
	categoryId := uint(uInt64)

	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid category ID"})
		return
	}

	address := ""
	if form.Value["address"] != nil {
		address = form.Value["address"][0]
	}

	description := ""
	if form.Value["description"] != nil {
		description = form.Value["description"][0]
	}
	// Create the document object
	docObj := models.Document{
		FileName:    file.Filename,
		FilePath:    utils.GetEnvVariable("BLOB_CONTAINER"),
		Description: description,
		CategoryID:  categoryId,
		Address:     address,
	}

	// upload the document
	document, err := contr.documentService.UploadDocument(userId, &docObj, openedFile)
	if err != nil {
		document = &models.Document{
			FileName: file.Filename,
			FilePath: "ERROR UPLOADING FILE",
		}
		c.JSON(400, gin.H{"message": "Failed to upload the files" + err.Error()})
	}

	// create the document tag records
	tagIdsStr := form.Value["tagIds"]

	if tagIdsStr != nil {
		// create the tag ids from the string it is a comma separated string
		tagIds := strings.Split(tagIdsStr[0], ",")

		// create the document tag records
		var documentTags []models.DocumentTag
		for _, tagIdStr := range tagIds {
			tagId, err := strconv.Atoi(tagIdStr)
			if err != nil {
				c.JSON(400, gin.H{"message": "Invalid tag ID"})
				return
			}
			documentTag := models.DocumentTag{
				DocumentID: document.ID,
				TagID:      uint(tagId),
			}
			documentTags = append(documentTags, documentTag)
		}
		_, err = contr.documentTag.CreateDocumentTagsBulk(documentTags)
		if err != nil {
			c.JSON(400, gin.H{"message": "Failed to create document tags"})
		}
	}
	// Return the response
	c.JSON(200, gin.H{
		"message":  "Files uploaded successfully.",
		"document": document,
	})
}

// DownloadDocument downloads a document
// @Summary Download document
// @Description Download a document by its file location
// @Tags documents
// @Produce  application/octet-stream
// @Param fileLocation query string true "File Location"
// @Success 200 {file} file "Successfully downloaded the document"
// @Failure 400 {object} map[string]interface{} "Error: Bad Request"
// @Router /document/download [get]
func (contr *DocumentController) DownloadDocument(c *gin.Context) {
	// get the documentId from the request
	documentIdStr := c.Param("documentId")
	// convert the id to an integer
	documentId, err := strconv.Atoi(documentIdStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid document ID. Error: " + err.Error(),
		})
		return
	}

	// call the service
	file, err := contr.documentService.DownloadDocument(documentId)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to download the file. Error: " + err.Error(),
		})
		return
	}
	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {
			log.Println(err)
		}
	}(file.(io.Closer))

	// Read the content into a byte array
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the content as a response
	c.Data(http.StatusOK, "application/octet-stream", buf.Bytes())
}

// SearchDocumentsHandler searches for documents with the given search parameters
// @Summary Search documents
// @Description Search for documents with the given search parameters
// @Tags documents
// @Accept  json
// @Produce  json
// @Param searchObj body search.Search true "Search object"
// @Param full query bool false "Set to true to include full tag details" default(false)
// @Success 200 {object} []models.Document "Successfully retrieved the document"
// @Failure 400 {object} map[string]interface{} "Error: Bad Request"
// @Router /document/search [post]
func (contr *DocumentController) SearchDocumentsHandler(c *gin.Context) {
	// create a new search object
	var searchObj search.Search
	if err := c.ShouldBindJSON(&searchObj); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid search object " + err.Error(),
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
	documents, err := contr.documentService.SearchDocuments(searchObj, full)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No documents found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, documents)
}
