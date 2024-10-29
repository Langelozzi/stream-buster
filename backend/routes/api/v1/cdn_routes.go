package v1

import (
	"github.com/STREAM-BUSTER/stream-buster/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetCDNRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitCDNDependencies()

	group := router.Group("/cdn")
	{
		group.GET("/movie/:tmdbId", controller.GetMovieContent)
		group.GET("/tv/:tmdbId/:seasonNum/:episodeNum", controller.GetTVContent)
	}
}
