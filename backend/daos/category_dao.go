package daos

import (
	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/models/search"
	"github.com/DOC-SWAP/Docswap-backend/utils"
	"github.com/DOC-SWAP/Docswap-backend/utils/database"
)

type CategoryDao struct{}

func NewCategoryDao() *CategoryDao {
	return &CategoryDao{}
}

func (dao *CategoryDao) GetAllCategoriesDao() ([]models.Category, error) {
	db := database.GetInstance()

	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (dao *CategoryDao) GetCategoryDao(id int) (*models.Category, error) {
	db := database.GetInstance()

	var category models.Category
	if err := db.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (dao *CategoryDao) SearchCategoriesDao(searchObj search.Search) ([]models.Category, error) {
	db := database.GetInstance()

	var categories []models.Category

	query := utils.BuildSearchQuery(db, models.Category{}, searchObj)

	if err := query.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
