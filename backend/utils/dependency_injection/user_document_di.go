package dependency_injection

import (
	"github.com/DOC-SWAP/Docswap-backend/controllers"
	"github.com/DOC-SWAP/Docswap-backend/daos"
	daoInterfaces "github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/services"
	servInterfaces "github.com/DOC-SWAP/Docswap-backend/services/interfaces"
)

func InitUserDocumentDependencies() *controllers.UserDocumentController {
	var dao daoInterfaces.UserDocumentDaoInterface = daos.NewUserDocumentDao()
	var service servInterfaces.UserDocumentServiceInterface = services.NewUserDocumentService(dao)
	return controllers.NewUserDocumentController(service)
}
