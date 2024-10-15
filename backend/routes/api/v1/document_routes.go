package v1

import (
	"github.com/DOC-SWAP/Docswap-backend/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetDocumentRoutes(router *gin.RouterGroup) {
	documentController := dependency_injection.InitDocumentDependencies()

	documentGroup := router.Group("/document")
	{
		documentGroup.GET("/", documentController.GetAllDocumentsHandler)
		documentGroup.GET("/:documentId", documentController.GetDocumentHandler)

		documentGroup.POST("/", documentController.CreateDocumentHandler)
		documentGroup.POST("/bulk", documentController.CreateDocumentsBulkHandler)

		documentGroup.DELETE("/:documentId", documentController.DeleteDocumentHandler)
		documentGroup.DELETE("/bulk", documentController.DeleteDocumentsBulkHandler)

		documentGroup.POST("/search", documentController.SearchDocumentsHandler)

		documentGroup.POST("/upload", documentController.WriteSingleDocument)
		documentGroup.GET("/:documentId/download", documentController.DownloadDocument)
	}
}
