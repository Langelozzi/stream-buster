package services

import (
	"github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/models/search"
)

type CategoryService struct {
	dao interfaces.CategoryDaoInterface
}

func NewCategoryService(dao interfaces.CategoryDaoInterface) *CategoryService {
	return &CategoryService{dao: dao}
}

func (service *CategoryService) GetAllCategories() ([]models.Category, error) {
	return service.dao.GetAllCategoriesDao()
}

func (service *CategoryService) GetCategory(id int) (*models.Category, error) {
	return service.dao.GetCategoryDao(id)
}
func (service *CategoryService) SearchCategories(searchObj search.Search) ([]models.Category, error) {
	return service.dao.SearchCategoriesDao(searchObj)
}
