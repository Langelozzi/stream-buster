package test

import (
	"fmt"
	"testing"

	"github.com/STREAM-BUSTER/stream-buster/daos"
	"github.com/stretchr/testify/assert"
)

func TestMediaDao(t *testing.T) {

	mediaDao := daos.NewMediaDao()

	t.Run("Get a media record from the dao", func(t *testing.T) {
		response, err := mediaDao.GetMediaById(1)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, response.ID, uint(1))
		assert.Equal(t, response.TMDBID, 1100)
		fmt.Println(response)
	})

}
