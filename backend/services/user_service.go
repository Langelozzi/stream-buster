package services

import (
	"github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/models"
)

type UserService struct {
	dao      interfaces.UserDaoInterface
	usageDao interfaces.UsageDaoInterface
}

func NewUserService(dao interfaces.UserDaoInterface, usageDao interfaces.UsageDaoInterface) *UserService {
	return &UserService{dao: dao, usageDao: usageDao}
}

func (service *UserService) GetAllUsers(includeDeleted bool, full bool) ([]models.User, error) {
	return service.dao.GetAllUsersDao(includeDeleted, full)
}

func (service *UserService) GetUser(id int, includeDeleted bool, full bool) (*models.User, error) {
	return service.dao.GetUserDao(id, includeDeleted, full)
}

func (service *UserService) GetUserByEmail(email string, includeDeleted bool, full bool) (*models.User, error) {
	return service.dao.GetUserByEmailDao(email, includeDeleted, full)
}

func (service *UserService) CreateUser(user *models.User) (*models.User, error) {
	return service.dao.CreateUserDao(user)
}

func (service *UserService) UpdateUser(updatedUser *models.User) (*models.User, error) {
	return service.dao.UpdateUserDao(updatedUser)
}

func (service *UserService) DeleteUser(id int, softDelete bool) error {
	if softDelete {
		return service.dao.SoftDeleteUserDao(id)
	}
	return service.dao.DeleteUserDao(id)
}

//func (service *UserService) GetUserUsageByUserId(userId int) (*models.Usage, error) {
//	return service.usageDao.GetUsageByUserId(userId)
//}
