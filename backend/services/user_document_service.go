package services

import (
	"github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/models"
)

type UserDocumentService struct {
	dao interfaces.UserDocumentDaoInterface
}

func NewUserDocumentService(dao interfaces.UserDocumentDaoInterface) *UserDocumentService {
	return &UserDocumentService{dao}
}

func (service *UserDocumentService) GetAllUserDocuments(includeDeleted bool, full bool) ([]models.UserDocument, error) {
	return service.dao.GetAllUserDocumentsDao(includeDeleted, full)
}

func (service *UserDocumentService) GetUserDocument(userId int, documentId int, includeDeleted bool, full bool) (*models.UserDocument, error) {
	return service.dao.GetUserDocumentDao(userId, documentId, includeDeleted, full)
}

func (service *UserDocumentService) CreateUserDocument(userDocument *models.UserDocument) (*models.UserDocument, error) {
	return service.dao.CreateUserDocumentDao(userDocument)
}

func (service *UserDocumentService) CreateUserDocumentsBulk(userDocuments []models.UserDocument) ([]models.UserDocument, error) {
	return service.dao.CreateUserDocumentsBulkDao(userDocuments)
}

func (service *UserDocumentService) DeleteUserDocument(userId int, documentId int, softDelete bool) error {
	if softDelete {
		return service.dao.SoftDeleteUserDocumentDao(userId, documentId)
	}
	return service.dao.DeleteUserDocumentDao(userId, documentId)
}

func (service *UserDocumentService) DeleteUserDocumentsBulk(userDocuments []models.UserDocument, softDelete bool) error {
	if softDelete {
		return service.dao.SoftDeleteUserDocumentsBulkDao(userDocuments)
	}
	return service.dao.DeleteUserDocumentsBulkDao(userDocuments)
}
