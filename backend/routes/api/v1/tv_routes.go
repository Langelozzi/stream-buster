package v1

import (
	"github.com/STREAM-BUSTER/stream-buster/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetTVRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitTVDependencies()

	group := router.Group("/tv")
	{
		group.GET("/:id", controller.GetTVDetails)
		group.GET("/:id/season/:seasonNum/episodes", controller.GetEpisodesInSeason)
	}
}
