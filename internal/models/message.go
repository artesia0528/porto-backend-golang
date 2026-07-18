package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Subject   string         `json:"subject"`
	Content   string         `json:"content" gorm:"type:text"`
	IsRead    bool           `json:"is_read" gorm:"default:false"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}