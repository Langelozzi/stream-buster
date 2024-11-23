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

// GetMovieContent retrieves the content for a movie by its tmdbId
// @Summary Get movie content by tmdbId
// @Description Retrieve HTML content for a movie based on the provided tmdbId
// @Tags cdn
// @Accept  json
// @Produce  html
// @Param tmdbId path string true "tmdbId of the movie"
// @Success 200 {string} string "HTML content of the movie"
// @Failure 400 {object} object "Error: Unable to procure content"
// @Router /cdn/movie/{tmdbId} [get]
func (contr *CDNController) GetMovieContent(c *gin.Context) {
	// get the tmdbId from the route params
	tmdbId := c.Param("tmdbId")

	// call the Service
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

// GetTVContent retrieves the content for a tv show by its tmdbId, season number and episode number
// @Summary Get movie content by tmdbId, season number and episode number
// @Description Retrieve HTML content for a tv episode based on the provided tmdbId, season number and episode number
// @Tags cdn
// @Accept  json
// @Produce  html
// @Param tmdbId path string true "tmdbId of the movie"
// @Param seasonNum path number true "season number of the episode"
// @Param episodeNum path number true "episode number of the episode"
// @Success 200 {string} string "HTML content of the movie"
// @Failure 400 {object} object "Error: Unable to procure content"
// @Router /cdn/tv/{tmdbId}/{seasonNum}/{episodeNum} [get]
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

	// call the Service
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
