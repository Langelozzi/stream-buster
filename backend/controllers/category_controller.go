package controllers

import (
	"fmt"
	"strconv"

	"github.com/DOC-SWAP/Docswap-backend/models/search"
	"github.com/DOC-SWAP/Docswap-backend/services/interfaces"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service interfaces.CategoryServiceInterface
}

func NewCategoryController(categoryService interfaces.CategoryServiceInterface) *CategoryController {
	return &CategoryController{
		service: categoryService,
	}
}

// GetAllCategoriesHandler GetAllCategories retrieves all categories
// @Summary Retrieve all categories
// @Description get all categories
// @Tags categories
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Category
// @Failure 400 {object} map[string]interface{} "Error: Bad Request"
// @Router /category [get]
func (contr *CategoryController) GetAllCategoriesHandler(c *gin.Context) {
	// call the service
	categories, err := contr.service.GetAllCategories()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No categories found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, categories)
}

// GetCategoryHandler GetCategory retrieves a category
// @Summary Retrieve a category
// @Description get a category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Failure 400 {object} map[string]interface{} "Error: Bad Request"
// @Router /category/{id} [get]
func (contr *CategoryController) GetCategoryHandler(c *gin.Context) {
	// get the id from the request
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Invalid id: %s. Error: %s", c.Param("id"), err.Error()),
		})
		return
	}

	// call the service
	category, err := contr.service.GetCategory(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Category record not found" + err.Error(),
		})
		return
	}

	c.JSON(200, category)
}

// SearchCategoriesHandler searches for categories with the given search parameters
// @Summary Search categories
// @Description Search for categories with the given search parameters
// @Tags categories
// @Accept  json
// @Produce  json
// @Param searchObj body search.Search true "Search object"
// @Success 200 {object} []models.Category "Successfully retrieved the category"
// @Failure 400 {object} map[string]interface{} "Error: Bad Request"
// @Router /category/search [post]
func (contr *CategoryController) SearchCategoriesHandler(c *gin.Context) {
	// create a new search object
	var searchObj search.Search
	if err := c.ShouldBindJSON(&searchObj); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid search object " + err.Error(),
		})
		return
	}

	// call the service
	categories, err := contr.service.SearchCategories(searchObj)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No categories found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, categories)
}
