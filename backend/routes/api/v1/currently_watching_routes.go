package v1

import (
	"github.com/STREAM-BUSTER/stream-buster/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetCurrentlyWatchingRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitCurrentlyWatchingDependencies()

	group := router.Group("/currently-watching")
	{
		group.POST("/", controller.CreateCurrentlyWatchingHandler)
		group.GET("/getall", controller.GetAllCurrentlyWatchingHandler)
		group.GET("/watchlist", controller.GetWatchlist)
		group.PUT("/update", controller.UpdateCurrentlyWatchingHandler)
		group.DELETE("/delete/:mediaId", controller.DeleteCurrentlyWatchingHandler)
	}
}
