package repositories

import (
	"portfolio-backend/internal/models"

	"gorm.io/gorm"
)

// ExperienceRepository menangani akses database untuk entity Experience.
type ExperienceRepository struct {
	db *gorm.DB
}

// NewExperienceRepository membuat instance baru ExperienceRepository.
func NewExperienceRepository(db *gorm.DB) *ExperienceRepository {
	return &ExperienceRepository{db: db}
}

// FindAll mengambil semua experience, diurutkan dari yang terbaru.
func (r *ExperienceRepository) FindAll() ([]models.Experience, error) {
	var experiences []models.Experience
	result := r.db.Order("start_date DESC").Find(&experiences)
	return experiences, result.Error
}

// FindByID mencari experience berdasarkan ID.
// Mengembalikan error jika tidak ditemukan.
func (r *ExperienceRepository) FindByID(id string) (*models.Experience, error) {
	var experience models.Experience
	result := r.db.Where("id = ?", id).First(&experience)
	if result.Error != nil {
		return nil, result.Error
	}
	return &experience, nil
}

// Create menyimpan experience baru ke database.
func (r *ExperienceRepository) Create(experience *models.Experience) error {
	return r.db.Create(experience).Error
}

// Update menyimpan perubahan experience ke database.
func (r *ExperienceRepository) Update(experience *models.Experience) error {
	return r.db.Save(experience).Error
}

// Delete menghapus experience berdasarkan ID.
// Mengembalikan jumlah rows yang terhapus.
func (r *ExperienceRepository) Delete(id string) (int64, error) {
	result := r.db.Where("id = ?", id).Delete(&models.Experience{})
	return result.RowsAffected, result.Error
}
