package database

import (
	"log/slog"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SeedAdmin membuat admin user pertama jika belum ada user di database.
func SeedAdmin(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count > 0 {
		return
	}

	hashedPassword, err := utils.HashPassword("admin123")
	if err != nil {
		slog.Error("Gagal hash password untuk admin seed", "error", err)
		return
	}

	admin := models.User{
		ID:       uuid.New().String(), // Generate random UUID
		Username: "admin",
		Password: hashedPassword,
	}

	if err := db.Create(&admin).Error; err != nil {
		slog.Error("Gagal membuat admin user", "error", err)
		return
	}

	slog.Info("Admin user berhasil dibuat", "username", "admin", "id", admin.ID)
}
