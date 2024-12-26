package controllers

import (
	"net/http"
	"strconv"

	"github.com/STREAM-BUSTER/stream-buster/models/db"
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

type MediaController struct {
	service interfaces.MediaServiceInterface
}

func NewMediaController(service interfaces.MediaServiceInterface) *MediaController {
	return &MediaController{
		service: service,
	}
}

// CreateMedia creates a new media record
// @Summary Create a new media record
// @Description create a new media record
// @Tags media
// @Accept  json
// @Produce  json
// @Param media body db.Media true "Media object that needs to be created"
// @Success 200 {object} db.Media "Successfully created the media record"
// @Failure 400 {object} map[string]interface{} "Error: Invalid request body or creation failed"
// @Router /media [post]
func (contr MediaController) CreateMedia(c *gin.Context) {
	media := &db.Media{}
	err := c.ShouldBindJSON(media)
	if err != nil {
		c.JSON(400, gin.H{
			"meesge": "Invalid request body. Error: " + err.Error(),
		})
		return
	}

	createdMedia, err := contr.service.CreateMedia(media)
	if pgError, ok := err.(*pgconn.PgError); ok {
		if pgError.Code == "23505" {
			createdMedia, err = contr.service.GetMediaByTMDBId(int64(media.TMDBID))
			if err != nil {
				c.String(http.StatusInternalServerError, "Record already exists error fetching")
			}
			c.JSON(http.StatusOK, createdMedia)
			return
		} else {
			c.JSON(400, gin.H{
				"message": "Failed to create a Media record. PostgreSQL Error Code: " + pgError.Code,
			})
			return
		}

	} else if err != nil {

		c.JSON(400, gin.H{
			"message": "Failed to create a currently watching record. Error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, createdMedia)

}

// GetMediaById retrieves a media record by its ID
// @Summary Retrieve a media record by ID
// @Description get a media record by its ID
// @Tags media
// @Accept  json
// @Produce  json
// @Param id query int true "Media ID"
// @Success 200 {object} db.Media "Successfully retrieved the media record"
// @Failure 400 {string} string "Error: Invalid media ID"
// @Router /media/by-id [get]
func (contr MediaController) GetMediaById(c *gin.Context) {
	mediaId, err := strconv.ParseInt(c.Query("id"), 10, 32)
	if err != nil {
		c.String(400, "Error parsing mediaId")
	}
	media, err := contr.service.GetMediaById(mediaId)
	c.JSON(200, media)
}

// GetMediaByTMDBId retrieves a media record by its TMDB ID
// @Summary Retrieve a media record by TMDB ID
// @Description get a media record by its TMDB ID
// @Tags media
// @Accept  json
// @Produce  json
// @Param tmdb_id query int true "TMDB ID"
// @Success 200 {object} db.Media "Successfully retrieved the media record"
// @Failure 400 {string} string "Error: Invalid TMDB ID"
// @Router /media/by-tmdb-id [get]
func (contr MediaController) GetMediaByTMDBId(c *gin.Context) {
	mediaId, err := strconv.ParseInt(c.Query("tmdb_id"), 10, 32)
	if err != nil {
		c.String(400, "Error parsing mediaId")
	}
	media, err := contr.service.GetMediaById(mediaId)
	c.JSON(200, media)
}
