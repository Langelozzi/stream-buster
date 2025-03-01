package controllers

import (
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/gin-gonic/gin"
	"strconv"
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
// @Param page query string true "The page to fetch. Pages are 20 items long."
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

	// get the page number
	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to get page number.",
		})
		return
	}

	// call the Service
	content, err := contr.service.SearchMultiMedia(query, int(page))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error getting search results. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, content)
}
