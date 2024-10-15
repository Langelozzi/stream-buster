package daos

import (
	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/models/search"
	"github.com/DOC-SWAP/Docswap-backend/utils"
	"github.com/DOC-SWAP/Docswap-backend/utils/database"
	"strings"
	"time"
)

type DocumentDao struct{}

func NewDocumentDao() *DocumentDao {
	return &DocumentDao{}
}

func (dao *DocumentDao) GetAllDocumentsDao(includeDeleted bool, full bool) ([]models.Document, error) {
	db := database.GetInstance()

	var documents []models.Document
	var query = db.Model(&models.Document{})

	if full {
		query = query.Preload("Tags").Preload("Category").Preload("ApprovedByUser").Preload("Tags.Tag")
	}

	if !includeDeleted {
		query = query.Where("deleted_at is null")
	}

	if err := query.Find(&documents).Error; err != nil {
		return nil, err // Return nil and the error if there's an issue with the query
	}

	return documents, nil
}

func (dao *DocumentDao) GetDocumentDao(documentId int, includeDeleted bool, full bool) (*models.Document, error) {
	db := database.GetInstance()
	var document models.Document
	var query = db.Model(&models.Document{})

	if full {
		query = query.Preload("Tags").Preload("Category").Preload("ApprovedByUser")
	}
	if !includeDeleted {
		query = query.Where("deleted_at IS NULL")
	}

	if err := query.First(&document, documentId).Error; err != nil {
		return nil, err
	}

	return &document, nil
}

func (dao *DocumentDao) CreateDocumentDao(document *models.Document) (*models.Document, error) {
	db := database.GetInstance()
	if err := db.Create(&document).Error; err != nil {
		return nil, err
	}
	return document, nil
}

func (dao *DocumentDao) CreateDocumentsBulkDao(documents []models.Document) ([]models.Document, error) {
	db := database.GetInstance()

	if err := db.Create(&documents).Error; err != nil {
		return nil, err
	}

	return documents, nil
}

func (dao *DocumentDao) DeleteDocumentDao(documentId int) error {
	db := database.GetInstance()

	if err := db.Delete(&models.Document{}, documentId).Error; err != nil {
		return err
	}

	return nil
}

func (dao *DocumentDao) DeleteDocumentsBulkDao(documents []models.Document) error {
	db := database.GetInstance()

	var conditions []string
	var params []interface{}

	for _, doc := range documents {
		conditions = append(conditions, "(id = ?)")
		params = append(params, doc.ID)
	}

	query := "DELETE FROM documents WHERE " + strings.Join(conditions, " OR ")

	if err := db.Exec(query, params...).Error; err != nil {
		return err
	}

	return nil
}

func (dao *DocumentDao) SoftDeleteDocumentDao(documentId int) error {
	db := database.GetInstance()

	if err := db.Model(&models.Document{}).Where("id = ?", documentId).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

func (dao *DocumentDao) SoftDeleteDocumentsBulkDao(documents []models.Document) error {
	db := database.GetInstance()

	var conditions []string
	var params []interface{}
	var now = time.Now()

	for _, doc := range documents {
		conditions = append(conditions, "(id = ?)")
		params = append(params, doc.ID)
	}

	query := "UPDATE documents SET deleted_at = ? WHERE " + strings.Join(conditions, " OR ")
	params = append([]interface{}{now}, params...)

	if err := db.Exec(query, params...).Error; err != nil {
		return err
	}

	return nil
}

func (dao *DocumentDao) SearchDocumentsDao(searchObj search.Search, full bool) ([]models.Document, error) {
	db := database.GetInstance()

	var documents []models.Document

	query := utils.BuildSearchQuery(db, models.Document{}, searchObj)

	if full {
		query = query.Preload("Tags").Preload("Category").Preload("ApprovedByUser").Preload("Tags.Tag")
	}

	if err := query.Find(&documents).Error; err != nil {
		return nil, err
	}

	return documents, nil
}
