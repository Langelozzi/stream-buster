package v1

import (
	"github.com/STREAM-BUSTER/stream-buster/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetSearchRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitSearchDependencies()

	group := router.Group("/search")
	{
		group.GET("/multi", controller.GetMultiMediaSearchResults)
	}
}
