package auth

import (
	"strconv"

	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) (models.User, error) {
	// Parse ID
	idStr := c.GetString("ID")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return models.User{}, err
	}

	// Create the user model
	user := models.User{
		ID:        id,
		Email:     c.GetString("Email"),
		FirstName: c.GetString("FirstName"),
		LastName:  c.GetString("LastName"),
	}

	return user, nil
}
