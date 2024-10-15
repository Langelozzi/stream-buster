package interfaces

import (
	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/models/search"
	"io"
)

type DocumentServiceInterface interface {
	GetAllDocuments(includeDeleted bool, full bool) ([]models.Document, error)
	GetDocument(documentId int, includeDeleted bool, full bool) (*models.Document, error)

	CreateDocument(document *models.Document) (*models.Document, error)
	CreateDocumentsBulk(documents []models.Document) ([]models.Document, error)

	DeleteDocument(documentId int, softDelete bool) error
	DeleteDocumentsBulk(documents []models.Document, softDelete bool) error

	SearchDocuments(searchObj search.Search, full bool) ([]models.Document, error)

	UploadDocument(userId uint, document *models.Document, file io.Reader) (*models.Document, error)
	DownloadDocument(documentId int) (io.Reader, error)
}
