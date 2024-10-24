package routes

import (
	"github.com/STREAM-BUSTER/stream-buster/daos"
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/middlewares"
	v1 "github.com/STREAM-BUSTER/stream-buster/routes/api/v1"
	"github.com/STREAM-BUSTER/stream-buster/services"
	iServices "github.com/STREAM-BUSTER/stream-buster/services/interfaces"
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
		v1.SetAuthRoutes(v1RouterGroup)
	}

	var userDao iDao.UserDaoInterface = daos.NewUserDao()
	var userService iServices.UserServiceInterface = services.NewUserService(userDao)
	var authDao iDao.AuthDaoInterface = daos.NewAuthDao()
	var authService iServices.AuthServiceInterface = services.NewAuthService(authDao, userService)

	// Setup private routes (requires authentication)
	privateRouterGroup := v1RouterGroup.Group("")
	privateRouterGroup.Use(middlewares.Auth(authService))
	{
		v1.SetUserRoutes(privateRouterGroup)
	}

	return router
}
