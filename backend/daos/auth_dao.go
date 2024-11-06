package daos

import (
	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/STREAM-BUSTER/stream-buster/utils/database"
)

type AuthDao struct{}

func NewAuthDao() *AuthDao {
	return &AuthDao{}
}

func (dao AuthDao) RegisterUser(user models.User) (*models.User, error) {
	db := database.GetInstance()

	// Create the user record
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	// Create a userRole record to make basic user
	userRole := models.UserRole{
		UserID: uint(user.ID),
		RoleID: 2,
	}
	if err := db.Create(&userRole).Error; err != nil {
		return nil, err
	}

	// Create a usage record for tracking usage
	usage := models.Usage{
		UserID:       uint(user.ID),
		RequestCount: 0,
	}
	if err := db.Create(&usage).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
