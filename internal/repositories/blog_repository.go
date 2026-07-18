package repositories

import (
	"portfolio-backend/internal/models"

	"gorm.io/gorm"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

func (r *BlogRepository) FindAll() ([]models.Blog, error) {
	var blogs []models.Blog
	result := r.db.Find(&blogs)
	return blogs, result.Error
}

func (r *BlogRepository) Create(blog *models.Blog) error {
	return r.db.Create(blog).Error
}

func (r *BlogRepository) FindByID(id string) (*models.Blog, error) {
	var blog models.Blog
	result := r.db.Where("id = ?",id).First(&blog)
	if result.Error != nil {
		return  nil, result.Error
	}
	return &blog, nil
}

func (r *BlogRepository) Update(blog *models.Blog) error {
	return r.db.Save(blog).Error
}

func (r *BlogRepository) Delete(id string) (int64, error) {
	result := r.db.Where("id = ?", id).Delete(&models.Blog{})
	return result.RowsAffected, result.Error
}