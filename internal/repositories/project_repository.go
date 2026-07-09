package repositories

import (
	"portfolio-backend/internal/models"

	"gorm.io/gorm"
)

// ProjectRepository menangani akses database untuk entity Project.
type ProjectRepository struct {
	db *gorm.DB
}

// NewProjectRepository membuat instance baru ProjectRepository.
func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

// FindAll mengambil semua project dari database.
func (r *ProjectRepository) FindAll() ([]models.Project, error) {
	var projects []models.Project
	result := r.db.Find(&projects)
	return projects, result.Error
}

// FindByID mencari project berdasarkan ID.
// Mengembalikan error jika tidak ditemukan.
func (r *ProjectRepository) FindByID(id uint) (*models.Project, error) {
	var project models.Project
	result := r.db.First(&project, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &project, nil
}

// Create menyimpan project baru ke database.
func (r *ProjectRepository) Create(project *models.Project) error {
	return r.db.Create(project).Error
}

// Update menyimpan perubahan project ke database.
func (r *ProjectRepository) Update(project *models.Project) error {
	return r.db.Save(project).Error
}

// Delete menghapus project berdasarkan ID.
// Mengembalikan jumlah rows yang terhapus.
func (r *ProjectRepository) Delete(id uint) (int64, error) {
	result := r.db.Delete(&models.Project{}, id)
	return result.RowsAffected, result.Error
}
