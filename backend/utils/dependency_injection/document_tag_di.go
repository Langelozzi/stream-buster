package dependency_injection

import (
	"github.com/DOC-SWAP/Docswap-backend/controllers"
	"github.com/DOC-SWAP/Docswap-backend/daos"
	daoInterfaces "github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/services"
	servInterfaces "github.com/DOC-SWAP/Docswap-backend/services/interfaces"
)

func InitDocumentTagDependencies() *controllers.DocumentTagController {
	var dao daoInterfaces.DocumentTagDaoInterface = daos.NewDocumentTagDao()
	var service servInterfaces.DocumentTagServiceInterface = services.NewDocumentTagService(dao)
	return controllers.NewDocumentTagController(service)
}
