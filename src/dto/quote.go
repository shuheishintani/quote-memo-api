package dto

import (
	"github.com/shuheishintani/quote-memo-api/src/models"
)

type QuoteInput struct {
	Text      string       `json:"text"`
	Page      int          `json:"page"`
	Published bool         `json:"published"`
	Book      models.Book  `json:"book"`
	Tags      []models.Tag `json:"tags"`
}
