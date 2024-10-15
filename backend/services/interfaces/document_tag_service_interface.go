package interfaces

import "github.com/DOC-SWAP/Docswap-backend/models"

type DocumentTagServiceInterface interface {
	GetAllDocumentTags(includeDeleted bool, full bool) ([]models.DocumentTag, error)
	GetDocumentTag(documentId int, tagId int, includeDeleted bool, full bool) (*models.DocumentTag, error)
	CreateDocumentTag(documentTag *models.DocumentTag) (*models.DocumentTag, error)
	CreateDocumentTagsBulk(documentTags []models.DocumentTag) ([]models.DocumentTag, error)
	DeleteDocumentTag(documentId int, tagId int, softDelete bool) error
	DeleteDocumentTagsBulk(documentTags []models.DocumentTag, softDelete bool) error
}
