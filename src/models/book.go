package models

import "time"

type Book struct {
	ID            int       `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	ISBN          string    `json:"isbn"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Publisher     string    `json:"publisher"`
	CoverImageUrl string    `json:"cover_image_url"`
	Quotes        []Quote   `json:"quotes"`
}
