package dependency_injection

import (
	"github.com/DOC-SWAP/Docswap-backend/controllers"
	"github.com/DOC-SWAP/Docswap-backend/daos"
	daoInterfaces "github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/services"
	servInterfaces "github.com/DOC-SWAP/Docswap-backend/services/interfaces"
)

func InitUserTagDependencies() *controllers.UserTagController {
	var dao daoInterfaces.UserTagDaoInterface = daos.NewUserTagDao()
	var service servInterfaces.UserTagServiceInterface = services.NewUserTagService(dao)
	return controllers.NewUserTagController(service)
}
