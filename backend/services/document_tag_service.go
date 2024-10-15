package services

import (
	"github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/models"
)

type DocumentTagService struct {
	dao interfaces.DocumentTagDaoInterface
}

func NewDocumentTagService(dao interfaces.DocumentTagDaoInterface) *DocumentTagService {
	return &DocumentTagService{dao}
}

func (service *DocumentTagService) GetAllDocumentTags(includeDeleted bool, full bool) ([]models.DocumentTag, error) {
	return service.dao.GetAllDocumentTagsDao(includeDeleted, full)
}

func (service *DocumentTagService) GetDocumentTag(documentId int, tagId int, includeDeleted bool, full bool) (*models.DocumentTag, error) {
	return service.dao.GetDocumentTagDao(documentId, tagId, includeDeleted, full)
}

func (service *DocumentTagService) CreateDocumentTag(documentTag *models.DocumentTag) (*models.DocumentTag, error) {
	return service.dao.CreateDocumentTagDao(documentTag)
}

func (service *DocumentTagService) CreateDocumentTagsBulk(documentTags []models.DocumentTag) ([]models.DocumentTag, error) {
	return service.dao.CreateDocumentTagsBulkDao(documentTags)
}

func (service *DocumentTagService) DeleteDocumentTag(documentId int, tagId int, softDelete bool) error {
	if softDelete {
		return service.dao.SoftDeleteDocumentTagDao(documentId, tagId)
	}
	return service.dao.DeleteDocumentTagDao(documentId, tagId)
}

func (service *DocumentTagService) DeleteDocumentTagsBulk(documentTags []models.DocumentTag, softDelete bool) error {
	if softDelete {
		return service.dao.SoftDeleteDocumentTagsBulkDao(documentTags)
	}
	return service.dao.DeleteDocumentTagsBulkDao(documentTags)
}
