package v1

import (
	"github.com/DOC-SWAP/Docswap-backend/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetUserDocumentRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitUserDocumentDependencies()

	documentGroup := router.Group("/userdocument")
	{
		documentGroup.GET("/", controller.GetAllUserDocumentsHandler)
		documentGroup.GET("/:userId/:documentId", controller.GetUserDocumentHandler)
		documentGroup.POST("/", controller.CreateUserDocumentHandler)
		documentGroup.POST("/bulk", controller.CreateUserDocumentsBulkHandler)
		documentGroup.DELETE("/:userId/:documentId", controller.DeleteUserDocumentHandler)
		documentGroup.DELETE("/bulk", controller.DeleteUserDocumentsBulkHandler)
	}
}
