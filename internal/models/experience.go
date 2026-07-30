package models

import (
	"time"

	"gorm.io/gorm"
)

// Experience merepresentasikan entity pengalaman kerja di database.
type Experience struct {
	ID          string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Company     string         `json:"company" gorm:"not null;size:100"`
	Position    string         `json:"position" gorm:"not null;size:100"`
	StartDate   string         `json:"start_date" gorm:"not null;size:10"`  // Format: "2024-01"
	EndDate     string         `json:"end_date" gorm:"size:10"`             // Kosong = masih bekerja
	IsCurrent   bool           `json:"is_current" gorm:"default:false"`
	Description string         `json:"description" gorm:"type:text"`
	LogoURL     string         `json:"logo_url" gorm:"size:500"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
