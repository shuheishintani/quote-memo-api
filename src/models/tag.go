package models

import "time"

type Tag struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at" validate:"datetime"`
	Name      string    `json:"name" validate:"required,min=1,max=100"`
}
