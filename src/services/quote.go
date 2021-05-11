package services

import (
	"github.com/shuheishintani/quote-manager-api/src/dto"
	"github.com/shuheishintani/quote-manager-api/src/models"
)

func (service *Service) GetPrivateQuotes(tagNames []string, uid string) ([]models.Quote, error) {
	quotes := []models.Quote{}

	if len(tagNames) == 0 {
		if result := service.db.Preload("Book").Preload("Tags").Where("UID = ?", uid).Find(&quotes); result.Error != nil {
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
		Where("UID = ?", uid).
		Find(&quotes); result.Error != nil {
		return []models.Quote{}, result.Error
	}
	return quotes, nil
}

func (service *Service) PostQuote(postQuoteInput dto.PostQuoteInput, uid string) (models.Quote, error) {
	book := models.Book{
		Isbn:          postQuoteInput.Book.Isbn,
		Title:         postQuoteInput.Book.Title,
		Author:        postQuoteInput.Book.Author,
		Publisher:     postQuoteInput.Book.Publisher,
		CoverImageUrl: postQuoteInput.Book.CoverImageUrl,
	}
	if result := service.db.FirstOrCreate(&book); result.Error != nil {
		return models.Quote{}, result.Error
	}

	quote := models.Quote{
		Text:   postQuoteInput.Text,
		Page:   postQuoteInput.Page,
		ISBN:   postQuoteInput.Book.Isbn,
		BookID: book.ID,
		UID:    uid,
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
