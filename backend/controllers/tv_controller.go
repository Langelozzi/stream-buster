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
