package controllers

import (
	"net/http"
	"strconv"

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

	err = contr.service.CreateMedia(media)

	c.String(http.StatusOK, "Media Created Successfully")

}
func (contr MediaController) GetMediaById(c *gin.Context) {
	mediaId, err := strconv.ParseInt(c.Query("id"), 10, 32)
	if err != nil {
		c.String(400, "Error parsing mediaId")
	}
	media, err := contr.service.GetMediaById(mediaId)
	c.JSON(200, media)
}
