package dto

import (
	"github.com/shuheishintani/quote-memo-api/src/models"
)

type QuoteInput struct {
	Text      string       `json:"text" validate:"required,min=1,max=400"`
	Page      int          `json:"page" validate:"gte=0,lte=50560"`
	Published bool         `json:"published" validate:"required"`
	Book      models.Book  `json:"book" validate:"required"`
	Tags      []models.Tag `json:"tags" validate:"required"`
}
