package dependency_injection

import (
	"github.com/DOC-SWAP/Docswap-backend/controllers"
	"github.com/DOC-SWAP/Docswap-backend/daos"
	daoInterfaces "github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/services"
	servInterfaces "github.com/DOC-SWAP/Docswap-backend/services/interfaces"
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
