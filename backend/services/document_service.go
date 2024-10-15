package services

import (
	"io"

	"github.com/DOC-SWAP/Docswap-backend/models/search"

	daosInterfaces "github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/models"
	servicesInterfaces "github.com/DOC-SWAP/Docswap-backend/services/interfaces"
)

type DocumentService struct {
	dao                 daosInterfaces.DocumentDaoInterface
	fileStorageService  servicesInterfaces.FileStorageServiceInterface
	userDocumentService servicesInterfaces.UserDocumentServiceInterface
}

func NewDocumentService(
	documentDao daosInterfaces.DocumentDaoInterface,
	fileStorageService servicesInterfaces.FileStorageServiceInterface,
	userDocumentService servicesInterfaces.UserDocumentServiceInterface) *DocumentService {
	return &DocumentService{
		dao:                 documentDao,
		fileStorageService:  fileStorageService,
		userDocumentService: userDocumentService,
	}
}

func (service *DocumentService) GetAllDocuments(includeDeleted bool, full bool) ([]models.Document, error) {
	return service.dao.GetAllDocumentsDao(includeDeleted, full)
}

func (service *DocumentService) GetDocument(documentId int, includeDeleted bool, full bool) (*models.Document, error) {
	return service.dao.GetDocumentDao(documentId, includeDeleted, full)
}

func (service *DocumentService) CreateDocument(document *models.Document) (*models.Document, error) {
	return service.dao.CreateDocumentDao(document)
}

func (service *DocumentService) CreateDocumentsBulk(documents []models.Document) ([]models.Document, error) {
	return service.dao.CreateDocumentsBulkDao(documents)
}

func (service *DocumentService) DeleteDocument(documentId int, softDelete bool) error {
	if softDelete {
		return service.dao.SoftDeleteDocumentDao(documentId)
	}
	return service.dao.DeleteDocumentDao(documentId)
}

func (service *DocumentService) DeleteDocumentsBulk(documents []models.Document, softDelete bool) error {
	if softDelete {
		return service.dao.SoftDeleteDocumentsBulkDao(documents)
	}
	return service.dao.DeleteDocumentsBulkDao(documents)
}

func (service *DocumentService) SearchDocuments(searchObj search.Search, full bool) ([]models.Document, error) {
	return service.dao.SearchDocumentsDao(searchObj, full)
}

func (service *DocumentService) UploadDocument(userId uint, document *models.Document, file io.Reader) (*models.Document, error) {
	// Upload file to file storage and get reference
	documentReference, err := service.fileStorageService.CreateFile(document, file)
	if err != nil {
		return nil, err
	}

	// Create document in database with reference
	document.FileStorageURL = documentReference
	createdDocument, err := service.CreateDocument(document)
	if err != nil {
		return nil, err
	}

	// Create UserDocument record for the current user and the document as owner
	userDocument := models.UserDocument{
		UserID:     userId,
		DocumentID: createdDocument.ID,
		IsOwner:    true,
	}
	_, err = service.userDocumentService.CreateUserDocument(&userDocument)
	if err != nil {
		return nil, err
	}

	return createdDocument, nil
}

func (service *DocumentService) DownloadDocument(documentId int) (io.Reader, error) {
	// Get the document object
	document, err := service.dao.GetDocumentDao(documentId, true, false)
	if err != nil {
		return nil, err
	}

	return service.fileStorageService.GetFile(document)
}
