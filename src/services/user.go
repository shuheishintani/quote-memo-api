package services

import (
	"fmt"

	"github.com/shuheishintani/quote-memo-api/src/dto"
	"github.com/shuheishintani/quote-memo-api/src/models"
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

func (service *Service) GetUserById(uid string) (models.User, error) {
	fmt.Println(uid)
	user := models.User{}
	if result := service.db.
		Preload("FavoriteQuotes.Tags").
		Preload("FavoriteQuotes.Book").
		Preload("FavoriteQuotes.User").
		Preload("Quotes", "published IS true").
		Preload("Quotes.Book").
		Preload("Quotes.Tags").
		First(&user, "id = ?", uid); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (service *Service) GetUsers() ([]models.User, error) {
	users := []models.User{}
	if result := service.db.Find(&users); result.Error != nil {
		return []models.User{}, result.Error
	}
	return users, nil
}
