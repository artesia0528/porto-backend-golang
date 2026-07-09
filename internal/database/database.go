package database

import (
	"log"
	"portfolio-backend/internal/config"
	"portfolio-backend/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open(config.AppConfig.DBPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek database:", err)
	}

	db.AutoMigrate(&models.User{}, &models.Project{})

	DB = db
}