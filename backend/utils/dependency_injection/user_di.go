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
	var usageDao daoInterfaces.UsageDaoInterface = daos.NewUsageDao()
	var service servInterfaces.UserServiceInterface = services.NewUserService(dao, usageDao)
	return controllers.NewUserController(service)
}
