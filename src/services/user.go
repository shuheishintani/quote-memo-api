package services

import (
	"github.com/shuheishintani/quote-manager-api/src/dto"
	"github.com/shuheishintani/quote-manager-api/src/models"
)

func (service *Service) CreateOrUpdateUser(userInput dto.UserInput) (models.User, error) {
	user := models.User{
		ID:              userInput.ID,
		Username:        userInput.Username,
		ProfileImageUrl: userInput.ProfileImageUrl,
		Provider:        userInput.Provider,
	}
	if result := service.db.Save(&user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
