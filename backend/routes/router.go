package routes

import (
	"github.com/STREAM-BUSTER/stream-buster/middlewares"
	v1 "github.com/STREAM-BUSTER/stream-buster/routes/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// Setup middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORS())

	// Setup routes for v1
	v1RouterGroup := router.Group("/api/v1")
	{
		v1.SetSearchRoutes(v1RouterGroup)
		v1.SetCDNRoutes(v1RouterGroup)
	}
	// {
	// 	// Setup public routes
	// 	publicRouterGroup := v1RouterGroup.Group("")
	// 	{
	// 		v1.SetAuthRoutes(publicRouterGroup)
	// 	}

	// 	// Setup private routes (requires authentication)
	// 	privateRouterGroup := v1RouterGroup.Group("")
	// 	privateRouterGroup.Use(middlewares.Auth())
	// 	{
	// 		v1.SetDocumentRoutes(privateRouterGroup)
	// 		v1.SetUserRoutes(privateRouterGroup)
	// 		v1.SetUserDocumentRoutes(privateRouterGroup)
	// 		v1.SetTagRoutes(privateRouterGroup)
	// 		v1.SetDocumentTagRoutes(privateRouterGroup)
	// 		v1.SetCategoryRoutes(privateRouterGroup)
	// 		v1.SetUserTagRoutes(v1RouterGroup)
	// 	}
	// }

	return router
}
