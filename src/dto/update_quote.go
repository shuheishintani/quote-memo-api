package dto

import "github.com/shuheishintani/quote-manager-api/src/models"

type UpdateQuoteInput struct {
	Text string       `json:"text"`
	Page int          `json:"page"`
	Tags []models.Tag `gorm:"many2many:quote_tags;" json:"tags"`
}
