package dependency_injection

import (
	"github.com/STREAM-BUSTER/stream-buster/controllers"
	"github.com/STREAM-BUSTER/stream-buster/daos"
	daoInterfaces "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/services"
	servInterfaces "github.com/STREAM-BUSTER/stream-buster/services/interfaces"
)

func InitUserDependencies() *controllers.UserController {
	var dao daoInterfaces.UserDaoInterface = daos.NewUserDao()
	var service servInterfaces.UserServiceInterface = services.NewUserService(dao)
	return controllers.NewUserController(service)
}

func InitUserServiceDependencies() servInterfaces.UserServiceInterface {
	var dao daoInterfaces.UserDaoInterface = daos.NewUserDao()
	return services.NewUserService(dao)
}
