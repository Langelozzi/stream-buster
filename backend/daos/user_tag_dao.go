package daos

import (
	"strings"
	"time"

	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/utils/database"
)

type UserTagDao struct{}

func NewUserTagDao() *UserTagDao {
	return &UserTagDao{}
}

func (dao *UserTagDao) GetAllUserTagsDao(includeDeleted bool, full bool) ([]models.UserTag, error) {
	db := database.GetInstance()

	var userTags []models.UserTag
	var query = db.Model(&models.UserTag{})

	if full {
		query = query.Preload("User").Preload("Tag")
	}
	if !includeDeleted {
		query = query.Where("deleted_at IS NULL")
	}

	if err := query.Find(&userTags).Error; err != nil {
		return nil, err
	}

	return userTags, nil
}

func (dao *UserTagDao) GetUserTagDao(userId int, tagId int, includeDeleted bool, full bool) (*models.UserTag, error) {
	db := database.GetInstance()

	var userTag models.UserTag
	var query = db.Model(&models.UserTag{})

	if full {
		query = query.Preload("User").Preload("Tag")
	}
	if !includeDeleted {
		query = query.Where("deleted_at IS NULL")
	}

	if err := query.First(&userTag, "user_id = ?", userId, "Tag_id = ?", tagId).Error; err != nil {
		return nil, err
	}

	return &userTag, nil
}

func (dao *UserTagDao) CreateUserTagDao(userTag *models.UserTag) (*models.UserTag, error) {
	db := database.GetInstance()

	if err := db.Create(&userTag).Error; err != nil {
		return nil, err
	}

	return userTag, nil
}

func (dao *UserTagDao) CreateUserTagsBulkDao(userTags []models.UserTag) ([]models.UserTag, error) {
	db := database.GetInstance()

	if err := db.Create(&userTags).Error; err != nil {
		return nil, err
	}

	return userTags, nil
}
func (dao *UserTagDao) DeleteUserTagDao(userId int, tagId int) error {
	db := database.GetInstance()

	if err := db.Where("user_id = ?", userId).Where("tag_id = ?", tagId).Delete(&models.UserTag{}).Error; err != nil {
		return err
	}

	return nil
}

func (dao *UserTagDao) DeleteUserTagsBulkDao(userTags []models.UserTag) error {

	db := database.GetInstance()

	var conditions []string
	var params []interface{}

	for _, doc := range userTags {
		conditions = append(conditions, "(user_id = ? AND tag_id = ?)")
		params = append(params, doc.UserID, doc.TagID)
	}

	query := "DELETE FROM user_tags WHERE " + strings.Join(conditions, " OR ")

	if err := db.Exec(query, params...).Error; err != nil {
		return err
	}

	return nil
}

func (dao *UserTagDao) SoftDeleteUserTagDao(userId int, documentId int) error {
	db := database.GetInstance()

	// set the deleted_at field to the current time
	if err := db.Model(&models.UserTag{}).Where("user_id = ?", userId).Where("tag_id = ?", documentId).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}
func (dao *UserTagDao) SoftDeleteUserTagsBulkDao(userTags []models.UserTag) error {
	db := database.GetInstance()

	var conditions []string
	var params []interface{}
	var now = time.Now()

	for _, doc := range userTags {
		conditions = append(conditions, "(user_id = ? AND tag_id = ?)")
		params = append(params, doc.UserID, doc.TagID)
	}

	query := "UPDATE user_tags SET deleted_at = ? WHERE " + strings.Join(conditions, " OR ")
	params = append([]interface{}{now}, params...)

	if err := db.Exec(query, params...).Error; err != nil {
		return err
	}

	return nil
}
