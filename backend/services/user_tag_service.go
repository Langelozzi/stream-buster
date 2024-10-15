package services

import (
	"github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/models"
)

type UserTagService struct {
	dao interfaces.UserTagDaoInterface
}

func NewUserTagService(dao interfaces.UserTagDaoInterface) *UserTagService {
	return &UserTagService{dao}
}

func (service *UserTagService) GetAllUserTags(includeDeleted bool, full bool) ([]models.UserTag, error) {
	return service.dao.GetAllUserTagsDao(includeDeleted, full)
}

func (service *UserTagService) GetUserTag(userId int, documentId int, includeDeleted bool, full bool) (*models.UserTag, error) {
	return service.dao.GetUserTagDao(userId, documentId, includeDeleted, full)
}

func (service *UserTagService) CreateUserTag(userTag *models.UserTag) (*models.UserTag, error) {
	return service.dao.CreateUserTagDao(userTag)
}

func (service *UserTagService) CreateUserTagsBulk(userTags []models.UserTag) ([]models.UserTag, error) {
	return service.dao.CreateUserTagsBulkDao(userTags)
}

func (service *UserTagService) DeleteUserTag(userId int, documentId int, softDelete bool) error {
	if softDelete {
		return service.dao.SoftDeleteUserTagDao(userId, documentId)
	}
	return service.dao.DeleteUserTagDao(userId, documentId)
}

func (service *UserTagService) DeleteUserTagsBulk(userTags []models.UserTag, softDelete bool) error {
	if softDelete {
		return service.dao.SoftDeleteUserTagsBulkDao(userTags)
	}
	return service.dao.DeleteUserTagsBulkDao(userTags)
}
