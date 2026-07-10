package models

import (
	"time"

	"gorm.io/gorm"
)

// User merepresentasikan entity user di database.
type User struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null;size:50"`
	Password  string         `json:"-" gorm:"not null"` // "-" = tidak pernah dikirim ke client
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
