package models

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Title     string         `json:"title"`
	Content   string         `json:"content" gorm:"type:text"`
	ImageURL  string         `json:"image_url" gorm:"size:500"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}