package interfaces

import "github.com/DOC-SWAP/Docswap-backend/models"

type UserDocumentDaoInterface interface {
	GetAllUserDocumentsDao(includeDeleted bool, full bool) ([]models.UserDocument, error)
	GetUserDocumentDao(userId int, documentId int, includeDeleted bool, full bool) (*models.UserDocument, error)
	CreateUserDocumentDao(userDocument *models.UserDocument) (*models.UserDocument, error)
	CreateUserDocumentsBulkDao(userDocuments []models.UserDocument) ([]models.UserDocument, error)
	DeleteUserDocumentDao(userId int, documentId int) error
	DeleteUserDocumentsBulkDao(userDocuments []models.UserDocument) error
	SoftDeleteUserDocumentDao(userId int, documentId int) error
	SoftDeleteUserDocumentsBulkDao(userDocuments []models.UserDocument) error
}
