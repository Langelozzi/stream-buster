package interfaces

import "github.com/DOC-SWAP/Docswap-backend/models"

type UserTagServiceInterface interface {
	GetAllUserTags(includeDeleted bool, full bool) ([]models.UserTag, error)
	GetUserTag(userId int, tagId int, includeDeleted bool, full bool) (*models.UserTag, error)
	CreateUserTag(userTag *models.UserTag) (*models.UserTag, error)
	CreateUserTagsBulk(userTags []models.UserTag) ([]models.UserTag, error)
	DeleteUserTag(userId int, tagId int, softDelete bool) error
	DeleteUserTagsBulk(userTags []models.UserTag, softDelete bool) error
}
