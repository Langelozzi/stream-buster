package v1

import (
	"github.com/DOC-SWAP/Docswap-backend/utils/dependency_injection"
	"github.com/gin-gonic/gin"
)

func SetCategoryRoutes(router *gin.RouterGroup) {
	categoryController := dependency_injection.InitCategoryDependencies()

	documentGroup := router.Group("/category")
	{
		documentGroup.GET("/", categoryController.GetAllCategoriesHandler)
		documentGroup.GET("/:id", categoryController.GetCategoryHandler)

		documentGroup.POST("/search", categoryController.SearchCategoriesHandler)
	}
}
