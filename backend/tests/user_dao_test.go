package test

import (
	"testing"

	"github.com/STREAM-BUSTER/stream-buster/daos"
	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/STREAM-BUSTER/stream-buster/utils/database"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserDao(t *testing.T) {
	db := database.GetInstance()

	// Ensure the test user does not already exist
	_ = db.Where("email = ?", "testuser@example.com").Delete(&models.User{})

	userDao := daos.UserDao{}

	t.Run("Successfully creating a user", func(t *testing.T) {
		user := &models.User{
			FirstName: "Test User",
			LastName:  "LastName",
			Email:     "testuser@example.com",
			Password:  "this is an encripted password",
		}

		createdUser, err := userDao.CreateUserDao(user)
		assert.NoError(t, err)
		assert.NotNil(t, createdUser)
		assert.Equal(t, user.Email, createdUser.Email)

		// Clean up the database after the test
		db.Where("email = ?", createdUser.Email).Delete(&models.User{})
	})

	t.Run("Try to create a user that already exists", func(t *testing.T) {
		user := &models.User{
			FirstName: "Test User",
			LastName:  "LastName",
			Email:     "testuser@example.com",
			Password:  "this is an encripted password",
		}
		// Create the user for the first time
		_, err := userDao.CreateUserDao(user)
		assert.NoError(t, err)

		anotherUser := &models.User{
			FirstName: "Test User",
			LastName:  "LastName",
			Email:     "testuser@example.com",
			Password:  "this is an encripted password",
		}
		// Attempt to create the same user again
		_, err = userDao.CreateUserDao(anotherUser)
		assert.Error(t, err) // Expecting an error here

		// Clean up the database after the test
		db.Where("email = ?", user.Email).Delete(&models.User{})
	})
}
