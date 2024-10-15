package daos

import (
	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/utils/database"
	"strings"
	"time"
)

type UserDocumentDao struct{}

func NewUserDocumentDao() *UserDocumentDao {
	return &UserDocumentDao{}
}

func (dao *UserDocumentDao) GetAllUserDocumentsDao(includeDeleted bool, full bool) ([]models.UserDocument, error) {
	db := database.GetInstance()

	var userDocuments []models.UserDocument
	var query = db.Model(&models.UserDocument{})

	if full {
		query = query.Preload("User").Preload("Document")
	}
	if !includeDeleted {
		query = query.Where("deleted_at IS NULL")
	}

	if err := query.Find(&userDocuments).Error; err != nil {
		return nil, err
	}

	return userDocuments, nil
}

func (dao *UserDocumentDao) GetUserDocumentDao(userId int, documentId int, includeDeleted bool, full bool) (*models.UserDocument, error) {
	db := database.GetInstance()

	var userDocument models.UserDocument
	var query = db.Model(&models.UserDocument{})

	if full {
		query = query.Preload("User").Preload("Document")
	}
	if !includeDeleted {
		query = query.Where("deleted_at IS NULL")
	}

	if err := query.First(&userDocument, "user_id = ?", userId, "document_id = ?", documentId).Error; err != nil {
		return nil, err
	}

	return &userDocument, nil
}

func (dao *UserDocumentDao) CreateUserDocumentDao(userDocument *models.UserDocument) (*models.UserDocument, error) {
	db := database.GetInstance()

	if err := db.Create(&userDocument).Error; err != nil {
		return nil, err
	}

	return userDocument, nil
}

func (dao *UserDocumentDao) CreateUserDocumentsBulkDao(userDocuments []models.UserDocument) ([]models.UserDocument, error) {
	db := database.GetInstance()

	if err := db.Create(&userDocuments).Error; err != nil {
		return nil, err
	}

	return userDocuments, nil
}

func (dao *UserDocumentDao) DeleteUserDocumentDao(userId int, documentId int) error {
	db := database.GetInstance()

	if err := db.Where("user_id = ?", userId).Where("document_id = ?", documentId).Delete(&models.UserDocument{}).Error; err != nil {
		return err
	}

	return nil
}

func (dao *UserDocumentDao) DeleteUserDocumentsBulkDao(userDocuments []models.UserDocument) error {
	db := database.GetInstance()

	var conditions []string
	var params []interface{}

	for _, doc := range userDocuments {
		conditions = append(conditions, "(user_id = ? AND document_id = ?)")
		params = append(params, doc.UserID, doc.DocumentID)
	}

	query := "DELETE FROM user_documents WHERE " + strings.Join(conditions, " OR ")

	if err := db.Exec(query, params...).Error; err != nil {
		return err
	}

	return nil
}

func (dao *UserDocumentDao) SoftDeleteUserDocumentDao(userId int, documentId int) error {
	db := database.GetInstance()

	// set the deleted_at field to the current time
	if err := db.Model(&models.UserDocument{}).Where("user_id = ?", userId).Where("document_id = ?", documentId).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

func (dao *UserDocumentDao) SoftDeleteUserDocumentsBulkDao(userDocuments []models.UserDocument) error {
	db := database.GetInstance()

	var conditions []string
	var params []interface{}
	var now = time.Now()

	for _, doc := range userDocuments {
		conditions = append(conditions, "(user_id = ? AND document_id = ?)")
		params = append(params, doc.UserID, doc.DocumentID)
	}

	query := "UPDATE user_documents SET deleted_at = ? WHERE " + strings.Join(conditions, " OR ")
	params = append([]interface{}{now}, params...)

	if err := db.Exec(query, params...).Error; err != nil {
		return err
	}

	return nil
}
