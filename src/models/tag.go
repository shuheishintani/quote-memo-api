package models

import "time"

type Tag struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}
