package interfaces

import "github.com/DOC-SWAP/Docswap-backend/models"

type UserTagDaoInterface interface {
	GetAllUserTagsDao(includeDeleted bool, full bool) ([]models.UserTag, error)
	GetUserTagDao(userId int, documentId int, includeDeleted bool, full bool) (*models.UserTag, error)
	CreateUserTagDao(userTag *models.UserTag) (*models.UserTag, error)
	CreateUserTagsBulkDao(userTags []models.UserTag) ([]models.UserTag, error)
	DeleteUserTagDao(userId int, documentId int) error
	DeleteUserTagsBulkDao(userTags []models.UserTag) error
	SoftDeleteUserTagDao(userId int, tagId int) error
	SoftDeleteUserTagsBulkDao(userTags []models.UserTag) error
}
