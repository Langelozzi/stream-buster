package interfaces

import "github.com/DOC-SWAP/Docswap-backend/models"

type UserServiceInterface interface {
	GetAllUsers(includeDeleted bool, full bool) ([]models.User, error)
	GetUser(id int, includeDeleted bool, full bool) (*models.User, error)
	GetUserByExternalId(externalId string, includeDeleted bool, full bool) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(updatedUser *models.User) (*models.User, error)
	UpdateUserByExternalID(externalId string, user *models.User) (*models.User, error)
	DeleteUser(id int, softDelete bool) error
}
