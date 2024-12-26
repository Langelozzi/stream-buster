package controllers

import (
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TVController struct {
	service interfaces.TVServiceInterface
}

func NewTVController(service interfaces.TVServiceInterface) *TVController {
	return &TVController{
		service: service,
	}
}

// GetTVDetails retrieves the details of a tv show by its tmdbId (id)
// @Summary Get tv show details by id
// @Description Retrieve the details of a tv show from TMDB using the id
// @Tags tv
// @Accept  json
// @Produce  json
// @Param id path string true "tmdbId of the tv show"
// @Success 200 {object} api.TV "The tv record"
// @Failure 400 {object} object "Error: Unable to procure content"
// @Router /tv/{id} [get]
func (contr *TVController) GetTVDetails(c *gin.Context) {
	// Get the series ID
	seriesId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No series ID provided. Error: " + err.Error(),
		})
		return
	}

	// call the Service
	content, err := contr.service.GetTVDetails(seriesId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to find details for tv series. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, content)
}

// GetEpisodesInSeason retrieves the list of episodes in a specific season of a tv show
// @Summary Get episodes of a show by season number
// @Description Retrieve a list of episodes in a specific season of a tv show
// @Tags tv
// @Accept  json
// @Produce  json
// @Param id path string true "tmdbId of the tv show"
// @Param seasonNum path number true "the season number"
// @Success 200 {array} api.TV "The list of tv records in that season"
// @Failure 400 {object} object "Error: Unable to procure content"
// @Router /tv/{id}/season/{seasonNum}/episodes [get]
func (contr *TVController) GetEpisodesInSeason(c *gin.Context) {
	// Get the series ID
	seriesId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No series ID provided. Error: " + err.Error(),
		})
		return
	}

	// Get the season number
	seasonNum, err := strconv.Atoi(c.Param("seasonNum"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No series ID provided. Error: " + err.Error(),
		})
		return
	}

	// call the Service
	content, err := contr.service.GetEpisodesInSeason(seriesId, seasonNum)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to find episodes for that season. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, content)
}
