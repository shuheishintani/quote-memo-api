package services

import (
	"fmt"

	"github.com/shuheishintani/quote-memo-api/src/models"
	"gorm.io/gorm/clause"
)

func (service *Service) CreateOrUpdateUser(userInput models.User) (models.User, error) {
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

func (service *Service) DeleteUser(uid string) (bool, error) {
	quotes, err := service.GetPrivateQuotes([]string{}, uid, 0, 10000)
	if err != nil {
		return false, err
	}

	if result := service.db.Select(clause.Associations).Delete(&quotes); result.Error != nil {
		return false, result.Error
	}

	if result := service.db.Select(clause.Associations).Delete(&models.User{ID: uid}); result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
