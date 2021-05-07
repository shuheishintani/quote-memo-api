package services

import "github.com/shuheishintani/quote-manager-api/src/models"

func (service *Service) GetTags() ([]models.Tag, error) {
	tags := []models.Tag{}
	if result := service.db.Find(&tags); result.Error != nil {
		return []models.Tag{}, result.Error
	}
	return tags, nil
}
