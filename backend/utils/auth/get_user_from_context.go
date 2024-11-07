package auth

import (
	"errors"
	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) (models.User, error) {
	// Parse ID
	idFloat, exists := c.Get("ID")
	if !exists {
		return models.User{}, errors.New("user ID not found in claims")
	}

	id := uint64(idFloat.(float64))

	// Create the user model
	user := models.User{
		ID:        id,
		Email:     c.GetString("Email"),
		FirstName: c.GetString("FirstName"),
		LastName:  c.GetString("LastName"),
	}

	return user, nil
}
