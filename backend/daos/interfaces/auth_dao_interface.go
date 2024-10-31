package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models"

type AuthDaoInterface interface {
	RegisterUser(user models.User) (*models.User, error)
}
