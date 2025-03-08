package v1

import (
	"github.com/STREAM-BUSTER/stream-buster/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetEndpointRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitEndpointDependencies()

	group := router.Group("/endpoint")
	{
		group.GET("", controller.GetAllEndpoints)
	}
}
