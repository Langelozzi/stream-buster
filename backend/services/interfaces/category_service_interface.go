package interfaces

import (
	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/models/search"
)

type CategoryServiceInterface interface {
	GetAllCategories() ([]models.Category, error)
	GetCategory(categoryId int) (*models.Category, error)
	SearchCategories(searchObj search.Search) ([]models.Category, error)
}
