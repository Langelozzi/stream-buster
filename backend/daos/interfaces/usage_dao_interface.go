package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models"

type UsageDaoInterface interface {
	GetUsageByUserId(userId int) ([]models.UserEndpointUsage, error)
}
