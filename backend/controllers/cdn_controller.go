package controllers

import (
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CDNController struct {
	service interfaces.CDNServiceInterface
}

func NewCDNController(service interfaces.CDNServiceInterface) *CDNController {
	return &CDNController{
		service: service,
	}
}

// GetMultiMediaSearchResults retrieves multimedia search results based on a query.
// @Summary Retrieve multimedia search results
// @Description Get multimedia content based on the search query.
// @Tags media
// @Accept  json
// @Produce  json
// @Param query query string true "Search query for multimedia content"
// @Success 200 {object} []interface{} "Successfully retrieved multimedia search results"
// @Failure 400 {object} map[string]interface{} "Error: Invalid or empty query, or no results found"
// @Router /media/search [get]
func (contr *CDNController) GetMovieContent(c *gin.Context) {
	// get the tmdbId from the route params
	tmdbId := c.Param("tmdbId")

	// call the service
	html, err := contr.service.GetMovieContent(tmdbId)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Unable to procure content for id: " + tmdbId,
			"error":   err,
		})
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

// GetMultiMediaSearchResults retrieves multimedia search results based on a query.
// @Summary Retrieve multimedia search results
// @Description Get multimedia content based on the search query.
// @Tags media
// @Accept  json
// @Produce  json
// @Param query query string true "Search query for multimedia content"
// @Success 200 {object} []interface{} "Successfully retrieved multimedia search results"
// @Failure 400 {object} map[string]interface{} "Error: Invalid or empty query, or no results found"
// @Router /media/search [get]
func (contr *CDNController) GetTVContent(c *gin.Context) {
	// get the tmdbId from the route params
	tmdbId := c.Param("tmdbId")

	// get the season num and convert to int
	seasonNum, err := strconv.Atoi(c.Param("seasonNum"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "seasonNum must be an integer",
		})
		return
	}

	// get the episode num and convert to int
	episodeNum, err := strconv.Atoi(c.Param("episodeNum"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "episodeNum must be an integer",
		})
		return
	}

	// call the service
	html, err := contr.service.GetTVContent(tmdbId, seasonNum, episodeNum)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Unable to procure content for id: " + tmdbId,
			"error":   err,
		})
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}
