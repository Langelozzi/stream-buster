package test

import (
	"testing"
	"time"

	"github.com/STREAM-BUSTER/stream-buster/daos"
	"github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/models/db"
	"github.com/STREAM-BUSTER/stream-buster/utils/database"
	"github.com/stretchr/testify/assert"
)

func TestCurrentlyWatchingDao(t *testing.T) {
	databaseInstance := database.GetInstance()

	var dao interfaces.CurrentlyWatchingDaoInterface = daos.NewCurrentlyWatchingDao()

	// Clean up any existing records for test consistency
	_ = databaseInstance.Where("user_id = ? AND media_id = ?", 1, 1).Delete(&db.CurrentlyWatching{})

	t.Run("Successfully creating a currently watching record", func(t *testing.T) {
		currentlyWatching := &db.CurrentlyWatching{
			UserID:        1,
			MediaID:       1,
			EpisodeNumber: 1,
			SeasonNumber:  1,
			CreatedAt:     pointerToTime(time.Now()),
		}

		createdRecord, err := dao.CreateCurrentlyWatching(currentlyWatching)
		assert.NoError(t, err)
		assert.NotNil(t, createdRecord)
		assert.Equal(t, currentlyWatching.UserID, createdRecord.UserID)
		assert.Equal(t, currentlyWatching.MediaID, createdRecord.MediaID)

		// Clean up the database after the test
		databaseInstance.Where("user_id = ? AND media_id = ?", currentlyWatching.UserID, currentlyWatching.MediaID).Delete(&db.CurrentlyWatching{})
	})

	t.Run("Attempt to create a duplicate currently watching record", func(t *testing.T) {
		currentlyWatching := &db.CurrentlyWatching{
			UserID:        1,
			MediaID:       1,
			EpisodeNumber: 1,
			SeasonNumber:  1,
			CreatedAt:     pointerToTime(time.Now()),
		}
		// Create the record for the first time
		_, err := dao.CreateCurrentlyWatching(currentlyWatching)
		assert.NoError(t, err)

		anotherWatching := &db.CurrentlyWatching{
			UserID:        1,
			MediaID:       1,
			EpisodeNumber: 1,
			SeasonNumber:  1,
		}
		// Attempt to create the same record again
		_, err = dao.CreateCurrentlyWatching(anotherWatching)
		assert.Error(t, err) // Expecting an error here

		// Clean up the database after the test
		databaseInstance.Where("user_id = ? AND media_id = ?", currentlyWatching.UserID, currentlyWatching.MediaID).Delete(&db.CurrentlyWatching{})
	})
}

func pointerToTime(t time.Time) *time.Time {
	return &t
}
