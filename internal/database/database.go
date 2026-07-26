package database

import (
	"fmt"
	"portfolio-backend/internal/config"
	"portfolio-backend/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Connect membuka koneksi database dan menjalankan auto-migration.
// Mengembalikan *gorm.DB dan error (bukan global variable).
func Connect(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gagal konek database: %w", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.Project{}, &models.Message{}, &models.Blog{}); err != nil {
		return nil, fmt.Errorf("gagal migrasi database: %w", err)
	}

	return db, nil
}