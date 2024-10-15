package v1

import (
	"github.com/DOC-SWAP/Docswap-backend/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetTagRoutes(router *gin.RouterGroup) {
	tagController := dependency_injection.InitTagDependencies()

	tagGroup := router.Group("/tag")
	{
		tagGroup.GET("/", tagController.GetAllTagsHandler)
		tagGroup.GET("/:tagId", tagController.GetTagHandler)
		tagGroup.POST("/", tagController.CreateTagHandler)
		tagGroup.POST("/bulk", tagController.CreateTagsBulkHandler)
		tagGroup.POST("/bulk/userTag", tagController.CreateTagsUserTagsBulkHandler)
		tagGroup.DELETE("/:tagId", tagController.DeleteTagHandler)
		tagGroup.DELETE("/bulk", tagController.DeleteTagsBulkHandler)
		tagGroup.POST("/search", tagController.SearchTagsHandler)
	}
}
