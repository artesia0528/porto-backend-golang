package models

import (
	"time"

	"gorm.io/gorm"
)

// Project merepresentasikan entity project portfolio di database.
type Project struct {
	ID          string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Title       string         `json:"title" gorm:"not null;size:100"`
	Description string         `json:"description" gorm:"type:text"`
	ImageURL    string         `json:"image_url" gorm:"size:500"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
