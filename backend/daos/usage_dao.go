package daos

import (
	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/STREAM-BUSTER/stream-buster/utils/database"
)

type UsageDao struct{}

func NewUsageDao() *UsageDao {
	return &UsageDao{}
}

func (dao *UsageDao) GetUsageByUserId(userId int) ([]models.UserEndpointUsage, error) {
	db := database.GetInstance()

	var usage []models.UserEndpointUsage
	query := db.Model(&models.UserEndpointUsage{})

	query.Preload("Endpoint")

	if err := query.Where("user_id = ?", userId).Find(&usage).Error; err != nil {
		return nil, err
	}

	return usage, nil
}
