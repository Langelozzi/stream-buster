package v1

import (
	"github.com/STREAM-BUSTER/stream-buster/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetMediaRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitMediaDependencies()

	group := router.Group("/media")
	{
		group.GET("", controller.GetMediaById)
		group.POST("/create", controller.CreateMedia)
	}
}
