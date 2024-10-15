package v1

import (
	"github.com/DOC-SWAP/Docswap-backend/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetUserTagRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitUserTagDependencies()

	UserTagGroup := router.Group("/usertag")
	{
		UserTagGroup.GET("/", controller.GetAllUserTagsHandler)
		UserTagGroup.GET("/:userId/:tagId", controller.GetUserTagHandler)
		UserTagGroup.POST("/", controller.CreateUserTagHandler)
		UserTagGroup.POST("/bulk", controller.CreateUserTagsBulkHandler)
		UserTagGroup.DELETE("/:userId/:tagId", controller.DeleteUserTagHandler)
		UserTagGroup.DELETE("/bulk", controller.DeleteUserTagsBulkHandler)
	}
}
