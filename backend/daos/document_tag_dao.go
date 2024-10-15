package daos

import (
	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/utils/database"
	"strings"
	"time"
)

type DocumentTagDao struct{}

func NewDocumentTagDao() *DocumentTagDao {
	return &DocumentTagDao{}
}

func (dao *DocumentTagDao) GetAllDocumentTagsDao(includeDeleted bool, full bool) ([]models.DocumentTag, error) {
	db := database.GetInstance()

	var documentTags []models.DocumentTag
	var query = db.Model(&models.DocumentTag{})

	if full {
		query = query.Preload("Document").Preload("Tag")
	}
	if !includeDeleted {
		query = query.Where("deleted_at IS NULL")
	}

	if err := query.Find(&documentTags).Error; err != nil {
		return nil, err
	}

	return documentTags, nil
}

func (dao *DocumentTagDao) GetDocumentTagDao(documentId int, tagId int, includeDeleted bool, full bool) (*models.DocumentTag, error) {
	db := database.GetInstance()

	var documentTag models.DocumentTag
	var query = db.Model(&models.DocumentTag{})

	if full {
		query = query.Preload("Document").Preload("Tag")
	}
	if !includeDeleted {
		query = query.Where("deleted_at IS NULL")
	}

	if err := query.First(&documentTag, "document_id = ?", documentId, "tag_id = ?", tagId).Error; err != nil {
		return nil, err
	}

	return &documentTag, nil
}

func (dao *DocumentTagDao) CreateDocumentTagDao(documentTag *models.DocumentTag) (*models.DocumentTag, error) {
	db := database.GetInstance()

	if err := db.Create(&documentTag).Error; err != nil {
		return nil, err
	}

	return documentTag, nil
}

func (dao *DocumentTagDao) CreateDocumentTagsBulkDao(documentTags []models.DocumentTag) ([]models.DocumentTag, error) {
	db := database.GetInstance()

	if err := db.Create(&documentTags).Error; err != nil {
		return nil, err
	}

	return documentTags, nil
}

func (dao *DocumentTagDao) DeleteDocumentTagDao(documentId int, tagId int) error {
	db := database.GetInstance()

	if err := db.Where("document_id = ?", documentId).Where("tag_id = ?", tagId).Delete(&models.DocumentTag{}).Error; err != nil {
		return err
	}

	return nil
}

func (dao *DocumentTagDao) DeleteDocumentTagsBulkDao(documentTags []models.DocumentTag) error {
	db := database.GetInstance()

	var conditions []string
	var params []interface{}

	for _, doc := range documentTags {
		conditions = append(conditions, "(document_id = ? AND tag_id = ?)")
		params = append(params, doc.DocumentID, doc.TagID)
	}

	query := "DELETE FROM document_tags WHERE " + strings.Join(conditions, " OR ")

	if err := db.Exec(query, params...).Error; err != nil {
		return err
	}

	return nil
}

func (dao *DocumentTagDao) SoftDeleteDocumentTagDao(documentId int, tagId int) error {
	db := database.GetInstance()

	// set the deleted_at field to the current time
	if err := db.Model(&models.DocumentTag{}).Where("document_id = ?", documentId).Where("tag_id = ?", tagId).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

func (dao *DocumentTagDao) SoftDeleteDocumentTagsBulkDao(documentTags []models.DocumentTag) error {
	db := database.GetInstance()

	var conditions []string
	var params []interface{}
	var now = time.Now()

	for _, doc := range documentTags {
		conditions = append(conditions, "(document_id = ? AND tag_id = ?)")
		params = append(params, doc.DocumentID, doc.TagID)
	}

	query := "UPDATE document_tags SET deleted_at = ? WHERE " + strings.Join(conditions, " OR ")
	params = append([]interface{}{now}, params...)

	if err := db.Exec(query, params...).Error; err != nil {
		return err
	}

	return nil
}
