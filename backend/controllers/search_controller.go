package controllers

import (
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/gin-gonic/gin"
)

type SearchController struct {
	service interfaces.SearchServiceInterface
}

func NewSearchController(service interfaces.SearchServiceInterface) *SearchController {
	return &SearchController{
		service: service,
	}
}

// GetMultiMediaSearchResults retrieves multimedia search results based on a query.
// @Summary Retrieve multimedia search results
// @Description Get multimedia content based on the search query.
// @Tags search
// @Accept  json
// @Produce  json
// @Param query query string true "Search query for multimedia content"
// @Success 200 {object} []interface{} "Successfully retrieved multimedia search results"
// @Failure 400 {object} map[string]interface{} "Error: Invalid or empty query, or no results found"
// @Router /search/multi [get]
func (contr *SearchController) GetMultiMediaSearchResults(c *gin.Context) {
	// get the query
	query := c.DefaultQuery("query", "")
	if len(query) == 0 {
		c.JSON(400, gin.H{
			"message": "Invalid or empty query.",
		})
		return
	}

	// call the Service
	content, err := contr.service.SearchMultiMedia(query)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No user records found. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, content)
}
