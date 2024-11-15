package v1

import (
	"github.com/STREAM-BUSTER/stream-buster/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetCurrentlyWatchingRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitCurrentlyWatchingDependencies()

	group := router.Group("/currently-watching")
	{
		group.GET("/getall", controller.GetAllCurrentlyWatchingHandler)
		group.GET("/watchlist", controller.GetWatchlist)
	}
}
