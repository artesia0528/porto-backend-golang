package repositories

import (
	"portfolio-backend/internal/models"

	"gorm.io/gorm"
)

// UserRepository menangani akses database untuk entity User.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository membuat instance baru UserRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByUsername mencari user berdasarkan username.
// Mengembalikan nil jika tidak ditemukan.
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Create menyimpan user baru ke database.
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// Count menghitung jumlah user di database.
func (r *UserRepository) Count() (int64, error) {
	var count int64
	result := r.db.Model(&models.User{}).Count(&count)
	return count, result.Error
}
