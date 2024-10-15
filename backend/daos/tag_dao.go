package daos

import (
	"gorm.io/gorm/clause"
	"strings"
	"time"

	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/models/search"
	"github.com/DOC-SWAP/Docswap-backend/utils"
	"github.com/DOC-SWAP/Docswap-backend/utils/database"
)

type TagDao struct{}

func NewTagDao() *TagDao {
	return &TagDao{}
}

func (dao *TagDao) GetAllTagsDao(includeDeleted bool, full bool) ([]models.Tag, error) {
	db := database.GetInstance()

	var tags []models.Tag
	var query = db.Model(&models.Tag{})

	if full {
		query = query.Preload("Document")
	}

	if !includeDeleted {
		query = query.Where("deleted_at is null")
	}

	if err := query.Find(&tags).Error; err != nil {
		return nil, err // Return nil and the error if there's an issue with the query
	}

	return tags, nil
}

func (dao *TagDao) GetTagDao(tagId int, includeDeleted bool, full bool) (*models.Tag, error) {
	db := database.GetInstance()
	var tag models.Tag
	var query = db.Model(&models.Tag{})

	if full {
		query = query.Preload("Document")
	}
	if !includeDeleted {
		query = query.Where("deleted_at IS NULL")
	}

	if err := query.First(&tag, tagId).Error; err != nil {
		return nil, err
	}

	return &tag, nil
}

func (dao *TagDao) CreateTagDao(tag *models.Tag) (*models.Tag, error) {
	db := database.GetInstance()
	if err := db.Create(&tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

func (dao *TagDao) CreateTagsBulkDao(tags []models.Tag) ([]models.Tag, error) {
	db := database.GetInstance()

	if err := db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (dao *TagDao) DeleteTagDao(tagId int) error {
	db := database.GetInstance()

	if err := db.Delete(&models.Tag{}, tagId).Error; err != nil {
		return err
	}

	return nil
}

func (dao *TagDao) DeleteTagsBulkDao(tags []models.Tag) error {
	db := database.GetInstance()

	var conditions []string
	var params []interface{}

	for _, tag := range tags {
		conditions = append(conditions, "(id = ?)")
		params = append(params, tag.ID)
	}

	query := "DELETE FROM tags WHERE " + strings.Join(conditions, " OR ")

	if err := db.Exec(query, params...).Error; err != nil {
		return err
	}

	return nil
}

func (dao *TagDao) SoftDeleteTagDao(tagId int) error {
	db := database.GetInstance()

	if err := db.Model(&models.Tag{}).Where("id = ?", tagId).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

func (dao *TagDao) SoftDeleteTagsBulkDao(tags []models.Tag) error {
	db := database.GetInstance()

	var conditions []string
	var params []interface{}
	var now = time.Now()

	for _, tag := range tags {
		conditions = append(conditions, "(id = ?)")
		params = append(params, tag.ID)
	}

	query := "UPDATE tags SET deleted_at = ? WHERE " + strings.Join(conditions, " OR ")
	params = append([]interface{}{now}, params...)

	if err := db.Exec(query, params...).Error; err != nil {
		return err
	}

	return nil
}

func (dao *TagDao) SearchTagsDao(searchObj search.Search) ([]models.Tag, error) {
	db := database.GetInstance()

	var tags []models.Tag

	query := utils.BuildSearchQuery(db, models.Tag{}, searchObj)

	if err := query.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}
