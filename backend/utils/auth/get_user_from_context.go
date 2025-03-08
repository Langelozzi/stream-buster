package auth

import (
	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) (models.User, error) {
	// Parse ID
	idFloat := c.GetFloat64("ID")
	id := uint64(idFloat)

	// Create the user model
	user := models.User{
		ID:        id,
		Email:     c.GetString("Email"),
		FirstName: c.GetString("FirstName"),
		LastName:  c.GetString("LastName"),
	}

	return user, nil
}
