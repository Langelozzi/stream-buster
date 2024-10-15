package dependency_injection

import (
	"github.com/DOC-SWAP/Docswap-backend/controllers"
	"github.com/DOC-SWAP/Docswap-backend/daos"
	daoInterfaces "github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/services"
	servInterfaces "github.com/DOC-SWAP/Docswap-backend/services/interfaces"
)

func InitCategoryDependencies() *controllers.CategoryController {
	var dao daoInterfaces.CategoryDaoInterface = daos.NewCategoryDao()
	var service servInterfaces.CategoryServiceInterface = services.NewCategoryService(dao)

	return controllers.NewCategoryController(service)
}
