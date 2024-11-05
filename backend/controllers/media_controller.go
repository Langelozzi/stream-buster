package controllers

import (
	"fmt"

	"github.com/STREAM-BUSTER/stream-buster/models/db"
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/gin-gonic/gin"
)

type MediaController struct {
	service interfaces.MediaServiceInterface
}

func NewMediaController(service interfaces.MediaServiceInterface) *MediaController {
	return &MediaController{
		service: service,
	}
}
func (contr MediaController) CreateMedia(c *gin.Context) {
	media := &db.Media{}
	err := c.ShouldBindJSON(media)
	if err != nil {
		c.JSON(400, gin.H{
			"meesge": "Invalid request body. Error: " + err.Error(),
		})
	}

	createdMedia, err := contr.service.CreateMedia(media)

	if createdMedia.ID != media.ID {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Failed to create a currently watching record. Expected ID %v, but got %v", createdMedia.ID, media.ID),
		})
		return
	}
	if createdMedia.TMDBID != media.TMDBID {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Expected TMDBID %v, but got %v", createdMedia.TMDBID, media.TMDBID),
		})
		return
	}
	if createdMedia.Title != media.Title {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Expected Title %v, but got %v", createdMedia.Title, media.Title),
		})
		return
	}
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create a currently watching record. Error: " + err.Error(),
		})
		return
	}

	c.JSON(201, createdMedia)

}
