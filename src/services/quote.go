package services

import (
	"fmt"

	"github.com/shuheishintani/quote-manager-api/src/dto"
	"github.com/shuheishintani/quote-manager-api/src/models"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (service *Service) GetQuotes(tagNames []string) ([]models.Quote, error) {
	fmt.Println(tagNames)
	quotes := []models.Quote{}

	if len(tagNames) == 0 {
		if result := service.db.Preload("Tags").Find(&quotes); result.Error != nil {
			return []models.Quote{}, result.Error
		}
		return quotes, nil
	}

	subQuery := service.db.
		Select("quote_id, count(*) AS count").
		Table("quote_tags qt").
		Joins("JOIN tags t ON qt.tag_id = t.id").
		Where("t.name IN (?)", tagNames).
		Group("quote_id")

	if result := service.db.Preload("Tags").
		Joins(
			"JOIN (?) AS matched ON quote_id = quotes.id AND matched.count = ?",
			subQuery,
			len(tagNames),
		).
		Find(&quotes); result.Error != nil {
		return []models.Quote{}, result.Error
	}

	return quotes, nil
}

func (service *Service) PostQuote(postQuoteInput dto.PostQuoteInput) (models.Quote, error) {
	quote := models.Quote{
		Text: postQuoteInput.Text,
		Page: postQuoteInput.Page,
		ISBN: postQuoteInput.ISBN,
	}

	if result := service.db.Save(&quote); result.Error != nil {
		return models.Quote{}, result.Error
	}

	for _, tag := range postQuoteInput.Tags {
		registeredTag := models.Tag{}
		if result := service.db.FirstOrCreate(&registeredTag, tag); result.Error != nil {
			return models.Quote{}, result.Error

		}
		if err := service.db.Model(&quote).Association("Tags").Append(&registeredTag); err != nil {
			return models.Quote{}, err
		}
	}
	return quote, nil
}
