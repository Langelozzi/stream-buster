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

// GetMovieSearchResults retrieves movie search results based on a query.
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
func (contr *SearchController) GetMovieSearchResults(c *gin.Context) {
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
	content, err := contr.service.SearchMovies(query, int(page))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error getting search results. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, content)
}

// GetTVSearchResults retrieves tv search results based on a query.
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
func (contr *SearchController) GetTVSearchResults(c *gin.Context) {
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
	content, err := contr.service.SearchTV(query, int(page))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error getting search results. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, content)
}

// GetTrendingMovies Get trending movies based on the time window.
// @Summary Retrieve trending movies
// @Description Get trending movies based on the specified time window.
// @Tags search
// @Accept  json
// @Produce  json
// @Param time_window query string true "Time window for trending movies (e.g., day, week)"
// @Param page query string true "The page to fetch. Pages are 20 items long."
// @Success 200 {object} []interface{} "Successfully retrieved trending movies"
// @Failure 400 {object} map[string]interface{} "Error: Invalid or empty time_window, or no results found"
// @Router /search/trending/movie [get]
func (contr *SearchController) GetTrendingMovies(c *gin.Context) {
	// get the timewindow
	timeWindow := c.DefaultQuery("time_window", "week")
	if len(timeWindow) == 0 {
		c.JSON(400, gin.H{
			"message": "Invalid or empty time_window.",
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
	content, err := contr.service.SearchTrendingMovies(timeWindow, int(page))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error getting search results. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, content)
}

// GetTrendingTV Get trending tv based on the time window.
// @Summary Retrieve trending tv
// @Description Get trending tv based on the specified time window.
// @Tags search
// @Accept  json
// @Produce  json
// @Param time_window query string true "Time window for trending tv (e.g., day, week)"
// @Param page query string true "The page to fetch. Pages are 20 items long."
// @Success 200 {object} []interface{} "Successfully retrieved trending tv"
// @Failure 400 {object} map[string]interface{} "Error: Invalid or empty time_window, or no results found"
// @Router /search/trending/tv [get]
func (contr *SearchController) GetTrendingTV(c *gin.Context) {
	// get the timewindow
	timeWindow := c.DefaultQuery("time_window", "week")
	if len(timeWindow) == 0 {
		c.JSON(400, gin.H{
			"message": "Invalid or empty time_window.",
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
	content, err := contr.service.SearchTrendingTV(timeWindow, int(page))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error getting search results. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, content)
}
