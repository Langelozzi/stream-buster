package daos

import (
	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/STREAM-BUSTER/stream-buster/utils/database"
)

type UsageDao struct{}

func NewUsageDao() *UsageDao {
	return &UsageDao{}
}

func (dao *UsageDao) GetUsageByUserId(userId int) (*models.Usage, error) {
	db := database.GetInstance()

	var usage models.Usage
	query := db.Model(&models.Usage{})

	if err := query.Where("user_id = ?", userId).First(&usage).Error; err != nil {
		return nil, err
	}

	return &usage, nil
}
