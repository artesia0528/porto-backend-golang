package services

import (
	"errors"
	"portfolio-backend/internal/dto"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"

	"github.com/google/uuid"
)

type BlogService struct {
	blogRepo *repositories.BlogRepository
}

func NewBlogService(blogRepo *repositories.BlogRepository) *BlogService {
	return &BlogService{blogRepo: blogRepo}
}

func (b *BlogService) GetAll() ([]models.Blog, error) {
	return b.blogRepo.FindAll()
}

func (b *BlogService) GetByID(id string) (*models.Blog, error) {
	blog, err := b.blogRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("Blog tidak ditemukan")
	}
	return blog, nil
}

func (b *BlogService) Create(req dto.CreateBlogRequest) (*models.Blog, error) {
	blog := &models.Blog{
		ID: uuid.New().String(),
		Title: req.Title,
		Content: req.Content,
		ImageURL: req.ImageURL,
	}

	if err := b.blogRepo.Create(blog); err != nil {
		return nil, errors.New("Gagal membuat blog")
	}
	return blog, nil
}

func (b *BlogService) Update(id string, req dto.UpdateBlogRequest) (*models.Blog, error) {
	blog, err := b.blogRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("Blog tidak ditemukan")
	}

	if req.Title != "" {
		blog.Title = req.Title
	}
	if req.Content != "" {
		blog.Content = req.Content
	}
	if req.ImageURL != "" {
		blog.ImageURL = req.ImageURL
	}

	if err := b.blogRepo.Update(blog); err != nil {
		return nil, errors.New("Gagal memperbarui blog")
	}

	return blog, nil
}

func (b *BlogService) Delete(id string) error {
	rowsAffected, err := b.blogRepo.Delete(id)
	if err != nil {
		return errors.New("Gagal menghapus blog")
	}
	if rowsAffected == 0 {
		return errors.New("Blog tidak ditemukan")
	}
	return nil
}