package interfaces

import (
	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/models/search"
)

type CategoryDaoInterface interface {
	GetAllCategoriesDao() ([]models.Category, error)
	GetCategoryDao(categoryId int) (*models.Category, error)
	SearchCategoriesDao(searchObj search.Search) ([]models.Category, error)
}
