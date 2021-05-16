package models

import "time"

type User struct {
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Username        string    `json:"username"`
	ProfileImageUrl string    `json:"profile_image_url"`
	Provider        string    `json:"provider"`
	Quotes          []Quote   `json:"quotes"`
}
