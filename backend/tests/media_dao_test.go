package test

import (
	"errors"
	"testing"

	"github.com/STREAM-BUSTER/stream-buster/daos"
	dbModels "github.com/STREAM-BUSTER/stream-buster/models/db"
	"github.com/STREAM-BUSTER/stream-buster/utils/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMediaDao(t *testing.T) {
	db := database.GetInstance()

	var media dbModels.Media
	// Ensure the test user does not already exist
	result := db.Where("id = ?", 2).First(&media)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.Create(&dbModels.Media{
			ID:          2,
			TMDBID:      1100,
			Title:       "How I met your mother",
			MediaTypeId: 1,
		})
	}

	mediaDao := daos.NewMediaDao()

	t.Run("Make sure that the dao returns a media object", func(t *testing.T) {
		media, err := mediaDao.GetMediaById(2, false)
		assert.Nil(t, err)
		assert.NotNil(t, media)
	})

	t.Run("Make sure that the correct object is being returned", func(t *testing.T) {
		media, err := mediaDao.GetMediaById(2, false)
		assert.Nil(t, err)
		correctMedia := dbModels.Media{
			ID:          2,
			TMDBID:      1100,
			Title:       "How I met your mother",
			MediaTypeId: 1,
		}
		// Check each field individually
		assert.Equal(t, correctMedia.ID, media.ID, "ID should match")
		assert.Equal(t, correctMedia.TMDBID, media.TMDBID, "TMDBID should match")
		assert.Equal(t, correctMedia.Title, media.Title, "Title should match")
		assert.Equal(t, correctMedia.MediaTypeId, media.MediaTypeId, "MediaTypeId should match")
	})
	t.Run("Make sure that the correct object is being returned", func(t *testing.T) {
		_, err := mediaDao.GetMediaById(10000, false)
		assert.NotNil(t, err)
	})

}
