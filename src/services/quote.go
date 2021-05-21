package services

import (
	"fmt"
	"strconv"

	"github.com/shuheishintani/quote-memo-api/src/dto"
	"github.com/shuheishintani/quote-memo-api/src/models"
	"gorm.io/gorm/clause"
)

func (service *Service) GetPublicQuotes(tagNames []string, offset int, limit int) ([]models.Quote, error) {
	quotes := []models.Quote{}

	if len(tagNames) == 0 {
		if result := service.db.
			Preload(clause.Associations).
			Where("published = true").
			Offset(offset).
			Limit(limit).
			Find(&quotes); result.Error != nil {
			return []models.Quote{}, result.Error
		}
		return quotes, nil
	}

	subQuery := service.db.
		Select("quote_id, count(*) AS count").
		Table("quotes_tags qt").
		Joins("JOIN tags t ON qt.tag_id = t.id").
		Where("t.name IN (?)", tagNames).
		Group("quote_id")

	if result := service.db.
		Preload(clause.Associations).
		Joins(
			"JOIN (?) AS matched ON quote_id = quotes.id AND matched.count = ?",
			subQuery,
			len(tagNames),
		).
		Where("published = true").
		Offset(offset).
		Limit(limit).
		Find(&quotes); result.Error != nil {
		return []models.Quote{}, result.Error
	}
	return quotes, nil

}

func (s *Service) GetQuoteById(id string) (models.Quote, error) {
	quote := models.Quote{}
	if result := s.db.Preload("Tags").Preload("Book").First(&quote, id); result.Error != nil {
		return models.Quote{}, result.Error
	}
	return quote, nil
}

func (service *Service) GetPrivateQuotes(tagNames []string, uid string, offset int, limit int) ([]models.Quote, error) {
	quotes := []models.Quote{}

	if len(tagNames) == 0 {
		if result := service.db.
			Preload("Tags").
			Preload("Book").
			Where("user_id = ?", uid).
			Order("created_at desc").
			Offset(offset).
			Limit(limit).
			Find(&quotes); result.Error != nil {
			return []models.Quote{}, result.Error
		}
		return quotes, nil
	}

	subQuery := service.db.
		Select("quote_id, count(*) AS count").
		Table("quotes_tags qt").
		Joins("JOIN tags t ON qt.tag_id = t.id").
		Where("t.name IN (?)", tagNames).
		Group("quote_id")

	if result := service.db.
		Preload(clause.Associations).
		Joins(
			"JOIN (?) AS matched ON quote_id = quotes.id AND matched.count = ?",
			subQuery,
			len(tagNames),
		).
		Where("user_id = ?", uid).
		Order("created_at desc").
		Offset(offset).
		Limit(limit).
		Find(&quotes); result.Error != nil {
		return []models.Quote{}, result.Error
	}
	return quotes, nil
}

func (service *Service) GetFavoriteQuotes(uid string) ([]models.Quote, error) {
	user, err := service.GetUserById(uid)
	if err != nil {
		return []models.Quote{}, err
	}
	return user.FavoriteQuotes, nil
}

func (service *Service) PostQuote(postQuoteInput dto.QuoteInput, uid string) (models.Quote, error) {
	book := models.Book{
		ISBN:          postQuoteInput.Book.ISBN,
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

func (service *Service) UpdateQuote(updateQuoteInput dto.QuoteInput, id string) (models.Quote, error) {
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

	if err := service.db.Model(&models.Quote{ID: i}).Association("FavoriteUsers").Clear(); err != nil {
		return false, err
	}

	if result := service.db.Select(clause.Associations).Delete(&models.Quote{ID: i}); result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (service *Service) AddFavoriteQuote(uid string, id string) (models.User, error) {
	quote, err := service.GetQuoteById(id)
	if err != nil {
		return models.User{}, err
	}

	fmt.Printf("%+v", quote)

	user := models.User{ID: uid}

	if err := service.db.Model(&user).Association("FavoriteQuotes").Append(&quote); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (service *Service) RemoveFavoriteQuote(uid string, id string) (models.User, error) {
	quote, err := service.GetQuoteById(id)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{ID: uid}

	if err := service.db.Model(&user).Association("FavoriteQuotes").Delete(&quote); err != nil {
		return models.User{}, err
	}
	return user, nil
}
