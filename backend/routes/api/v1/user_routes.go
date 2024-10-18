package v1

import (
	"github.com/STREAM-BUSTER/stream-buster/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetUserRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitUserDependencies()

	group := router.Group("/user")
	{
		group.GET("/", controller.GetAllUsersHandler)
		group.GET("/:id", controller.GetUserHandler)
		group.GET("/current", controller.GetCurrentUserHandler)
		group.POST("/", controller.CreateUserHandler)
		group.PUT("/:id", controller.UpdateUserHandler)
		//group.PUT("/current", controller.UpdateCurrentUserHandler)
		group.DELETE("/:id", controller.DeleteUserHandler)
	}
}
