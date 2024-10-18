package v1

import (
	"github.com/STREAM-BUSTER/stream-buster/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(router *gin.RouterGroup) {
	userController := dependency_injection.InitUserDependencies()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", userController.CreateUserHandler)
	}
}
