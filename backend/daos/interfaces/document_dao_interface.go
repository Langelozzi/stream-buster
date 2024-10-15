package interfaces

import (
	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/models/search"
)

type DocumentDaoInterface interface {
	GetAllDocumentsDao(includeDeleted bool, full bool) ([]models.Document, error)
	GetDocumentDao(documentId int, includeDeleted bool, full bool) (*models.Document, error)

	CreateDocumentDao(document *models.Document) (*models.Document, error)
	CreateDocumentsBulkDao(documents []models.Document) ([]models.Document, error)

	DeleteDocumentDao(documentId int) error
	DeleteDocumentsBulkDao(documents []models.Document) error

	SoftDeleteDocumentDao(documentId int) error
	SoftDeleteDocumentsBulkDao(documents []models.Document) error

	SearchDocumentsDao(searchObj search.Search, full bool) ([]models.Document, error)
}
