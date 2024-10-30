package v1

import (
	"github.com/STREAM-BUSTER/stream-buster/middlewares"
	"github.com/STREAM-BUSTER/stream-buster/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(router *gin.RouterGroup) {
	authController := dependency_injection.InitAuthDependencies()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authController.CreateUser)
		authGroup.POST("/login", authController.LoginUser)
		authGroup.POST("/logout", authController.LogoutUser)
		authGroup.GET("/test", middlewares.Auth(authController.Service), authController.TestAuthMiddleware)
	}
}
