package interfaces

import "github.com/DOC-SWAP/Docswap-backend/models"

type DocumentTagDaoInterface interface {
	GetAllDocumentTagsDao(includeDeleted bool, full bool) ([]models.DocumentTag, error)
	GetDocumentTagDao(documentId int, tagId int, includeDeleted bool, full bool) (*models.DocumentTag, error)
	CreateDocumentTagDao(documentTag *models.DocumentTag) (*models.DocumentTag, error)
	CreateDocumentTagsBulkDao(documentTags []models.DocumentTag) ([]models.DocumentTag, error)
	DeleteDocumentTagDao(documentId int, tagId int) error
	DeleteDocumentTagsBulkDao(documentTags []models.DocumentTag) error
	SoftDeleteDocumentTagDao(documentId int, tagId int) error
	SoftDeleteDocumentTagsBulkDao(documentTags []models.DocumentTag) error
}
