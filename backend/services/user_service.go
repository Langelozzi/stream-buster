package services

import (
	"github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/models"
)

type UserService struct {
	dao interfaces.UserDaoInterface
}

func NewUserService(dao interfaces.UserDaoInterface) *UserService {
	return &UserService{dao: dao}
}

func (service *UserService) GetAllUsers(includeDeleted bool, full bool) ([]models.User, error) {
	return service.dao.GetAllUsersDao(includeDeleted, full)
}

func (service *UserService) GetUser(id int, includeDeleted bool, full bool) (*models.User, error) {
	return service.dao.GetUserDao(id, includeDeleted, full)
}

func (service *UserService) GetUserByExternalId(externalId string, includeDeleted bool, full bool) (*models.User, error) {
	return service.dao.GetUserByExternalIdDao(externalId, includeDeleted, full)
}

func (service *UserService) CreateUser(user *models.User) (*models.User, error) {
	return service.dao.CreateUserDao(user)
}

func (service *UserService) UpdateUser(updatedUser *models.User) (*models.User, error) {
	return service.dao.UpdateUserDao(updatedUser)
}

func (service *UserService) UpdateUserByExternalID(externalID string, updatedData *models.User) (*models.User, error) {
	return service.dao.UpdateUserByExternalID(externalID, updatedData)
}

func (service *UserService) DeleteUser(id int, softDelete bool) error {
	if softDelete {
		return service.dao.SoftDeleteUserDao(id)
	}
	return service.dao.DeleteUserDao(id)
}
