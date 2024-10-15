package dependency_injection

import (
	"github.com/DOC-SWAP/Docswap-backend/controllers"
	"github.com/DOC-SWAP/Docswap-backend/daos"
	daoInterfaces "github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/services"
	servInterfaces "github.com/DOC-SWAP/Docswap-backend/services/interfaces"
)

func InitDocumentDependencies() *controllers.DocumentController {
	var azureBlobStorageDao daoInterfaces.FileStorageDaoInterface = daos.NewAzureBlobStorageDao()
	var azureBlobStorageService servInterfaces.FileStorageServiceInterface = services.NewAzureBlobStorageService(azureBlobStorageDao)

	var userDocumentDao daoInterfaces.UserDocumentDaoInterface = daos.NewUserDocumentDao()
	var userDocumentService servInterfaces.UserDocumentServiceInterface = services.NewUserDocumentService(userDocumentDao)

	var documentTagDao daoInterfaces.DocumentTagDaoInterface = daos.NewDocumentTagDao()
	var documentTagService servInterfaces.DocumentTagServiceInterface = services.NewDocumentTagService(documentTagDao)

	var documentDao daoInterfaces.DocumentDaoInterface = daos.NewDocumentDao()
	var documentService servInterfaces.DocumentServiceInterface = services.NewDocumentService(documentDao, azureBlobStorageService, userDocumentService)

	return controllers.NewDocumentController(documentService, documentTagService)
}
