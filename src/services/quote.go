package services

import (
	"strconv"

	"github.com/shuheishintani/quote-manager-api/src/dto"
	"github.com/shuheishintani/quote-manager-api/src/models"
	"gorm.io/gorm/clause"
)

func (service *Service) GetPublicQuotes() ([]models.Quote, error) {
	quotes := []models.Quote{}
	if result := service.db.Preload(clause.Associations).Where("Published = true").Find(&quotes); result.Error != nil {
		return []models.Quote{}, result.Error
	}
	return quotes, nil
}

func (s *Service) GetQuoteById(id string) (models.Quote, error) {
	quote := models.Quote{}
	if result := s.db.Preload("Book").Preload("Tags").First(&quote, id); result.Error != nil {
		return models.Quote{}, result.Error
	}
	return quote, nil
}

func (service *Service) GetPrivateQuotes(tagNames []string, uid string) ([]models.Quote, error) {
	quotes := []models.Quote{}

	if len(tagNames) == 0 {
		if result := service.db.Preload(clause.Associations).Where("user_id = ?", uid).Find(&quotes); result.Error != nil {
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

	if result := service.db.Preload("Book").Preload("Tags").
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
		ISBN:          postQuoteInput.Book.Isbn,
		Title:         postQuoteInput.Book.Title,
		Author:        postQuoteInput.Book.Author,
		Publisher:     postQuoteInput.Book.Publisher,
		CoverImageUrl: postQuoteInput.Book.CoverImageUrl,
	}
	if result := service.db.Where(book).FirstOrCreate(&book); result.Error != nil {
		return models.Quote{}, result.Error
	}

	quote := models.Quote{
		Text:      postQuoteInput.Text,
		Page:      postQuoteInput.Page,
		Published: postQuoteInput.Published,
		BookID:    book.ID,
		Book:      book,
		UserID:    uid,
	}
	if result := service.db.Save(&quote); result.Error != nil {
		return models.Quote{}, result.Error
	}

	for _, tag := range postQuoteInput.Tags {
		if result := service.db.Where(tag).FirstOrCreate(&tag); result.Error != nil {
			return models.Quote{}, result.Error
		}
		if err := service.db.Model(&quote).Association("Tags").Append(&tag); err != nil {
			return models.Quote{}, err
		}
	}
	return quote, nil
}

func (service *Service) UpdateQuote(updateQuoteInput dto.UpdateQuoteInput, id string) (models.Quote, error) {
	if result := service.db.Model(&models.Quote{}).Where("id = ?", id).Updates(map[string]interface{}{
		"Text":      updateQuoteInput.Text,
		"Page":      updateQuoteInput.Page,
		"Published": updateQuoteInput.Published,
	}); result.Error != nil {
		return models.Quote{}, result.Error
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		return models.Quote{}, err
	}

	for _, tag := range updateQuoteInput.Tags {
		if result := service.db.Where(tag).FirstOrCreate(&tag); result.Error != nil {
			return models.Quote{}, result.Error
		}

		if err := service.db.Model(&models.Quote{ID: i}).Association("Tags").Append(&tag); err != nil {
			return models.Quote{}, err
		}

	}

	updated, err := service.GetQuoteById(id)
	if err != nil {
		return models.Quote{}, err
	}
	return updated, nil
}

func (service *Service) DeleteQuote(id string) (result bool, err error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return false, err
	}
	if result := service.db.Select("Tags").Delete(&models.Quote{ID: i}); result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
