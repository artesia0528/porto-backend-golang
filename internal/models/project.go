package models

import "gorm.io/gorm"

// Project merepresentasikan entity project portfolio di database.
type Project struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null;size:100"`
	Description string `json:"description" gorm:"type:text"`
	ImageURL    string `json:"image_url" gorm:"size:500"`
}
