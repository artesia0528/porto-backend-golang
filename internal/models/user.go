package models

import "gorm.io/gorm"

// User merepresentasikan entity user di database.
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex;not null;size:50"`
	Password string `json:"-" gorm:"not null"` // "-" = tidak pernah dikirim ke client
}
