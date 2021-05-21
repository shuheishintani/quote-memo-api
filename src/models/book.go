package models

import "time"

type Book struct {
	ID            int       `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	ISBN          string    `json:"isbn" validate:"isbn,required"`
	Title         string    `json:"title" validate:"required,min=1,max=100"`
	Author        string    `json:"author" validate:"required,min=1,max=100"`
	Publisher     string    `json:"publisher" validate:"required,min=1,max=100"`
	CoverImageUrl string    `json:"cover_image_url" validate:"url"`
	Quotes        []Quote   `json:"quotes"`
}
