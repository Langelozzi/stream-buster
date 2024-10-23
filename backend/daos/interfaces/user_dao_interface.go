package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models"

type UserDaoInterface interface {
	GetAllUsersDao(includeDeleted bool, full bool) ([]models.User, error)
	GetUserDao(id int, includeDeleted bool, full bool) (*models.User, error)
	GetUserByUsernameDao(username string, includeDeleted bool, full bool) (*models.User, error)
	CreateUserDao(user *models.User) (*models.User, error)
	UpdateUserDao(updatedUser *models.User) (*models.User, error)
	DeleteUserDao(id int) error
	SoftDeleteUserDao(id int) error
}
