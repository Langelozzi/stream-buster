package v1

import (
	"github.com/DOC-SWAP/Docswap-backend/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetDocumentTagRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitDocumentTagDependencies()

	documentGroup := router.Group("/documenttag")
	{
		documentGroup.GET("/", controller.GetAllDocumentTagsHandler)
		documentGroup.GET("/:documentId/:tagId", controller.GetDocumentTagHandler)
		documentGroup.POST("/", controller.CreateDocumentTagHandler)
		documentGroup.POST("/bulk", controller.CreateDocumentTagsBulkHandler)
		documentGroup.DELETE("/:documentId/:tagId", controller.DeleteDocumentTagHandler)
		documentGroup.DELETE("/bulk", controller.DeleteDocumentTagsBulkHandler)
	}
}
