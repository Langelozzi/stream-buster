package interfaces

import "github.com/DOC-SWAP/Docswap-backend/models"

type UserDocumentServiceInterface interface {
	GetAllUserDocuments(includeDeleted bool, full bool) ([]models.UserDocument, error)
	GetUserDocument(userId int, documentId int, includeDeleted bool, full bool) (*models.UserDocument, error)
	CreateUserDocument(userDocument *models.UserDocument) (*models.UserDocument, error)
	CreateUserDocumentsBulk(userDocuments []models.UserDocument) ([]models.UserDocument, error)
	DeleteUserDocument(userId int, documentId int, softDelete bool) error
	DeleteUserDocumentsBulk(userDocuments []models.UserDocument, softDelete bool) error
}
