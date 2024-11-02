package v1

import (
	"github.com/STREAM-BUSTER/stream-buster/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetMovieRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitMovieDependencies()

	group := router.Group("/movie")
	{
		group.GET("/:id", controller.GetMovieDetails)
	}
}
