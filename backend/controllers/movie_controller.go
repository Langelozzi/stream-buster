package controllers

import (
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MovieController struct {
	service interfaces.MovieServiceInterface
}

func NewMovieController(service interfaces.MovieServiceInterface) *MovieController {
	return &MovieController{
		service: service,
	}
}

// GetMovieDetails retrieves the details of a movie by its tmdbId (id)
// @Summary Get movie details by id
// @Description Retrieve the details of a movie from TMDB using the id
// @Tags movie
// @Accept  json
// @Produce  json
// @Param id path string true "tmdbId of the movie"
// @Success 200 {object} api.Movie "The movie record"
// @Failure 400 {object} object "Error: Unable to procure content"
// @Router /movie/{id} [get]
func (contr *MovieController) GetMovieDetails(c *gin.Context) {
	// Get the series ID
	seriesId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No series ID provided. Error: " + err.Error(),
		})
		return
	}

	// call the Service
	content, err := contr.service.GetMovieDetails(seriesId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to find details for movie. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, content)
}
