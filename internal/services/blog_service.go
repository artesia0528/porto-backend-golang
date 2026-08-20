package services

import (
	"errors"
	"mime/multipart"
	"portfolio-backend/internal/dto"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/utils"

	"github.com/google/uuid"
)

// BlogService menangani business logic untuk blog.
type BlogService struct {
	blogRepo *repositories.BlogRepository
	baseURL  string
}

// NewBlogService membuat instance baru BlogService.
func NewBlogService(blogRepo *repositories.BlogRepository, baseURL string) *BlogService {
	return &BlogService{blogRepo: blogRepo, baseURL: baseURL}
}

// GetAll mengambil semua blog.
func (s *BlogService) GetAll() ([]models.Blog, error) {
	blogs, err := s.blogRepo.FindAll()
	if err != nil {
		return nil, err
	}

	for i := range blogs {
		blogs[i].ImageURL = s.buildFileURL(blogs[i].ImageURL)
	}

	return blogs, nil
}

// GetByID mengambil satu blog berdasarkan ID.
func (s *BlogService) GetByID(id string) (*models.Blog, error) {
	blog, err := s.blogRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("blog tidak ditemukan")
	}

	blog.ImageURL = s.buildFileURL(blog.ImageURL)
	return blog, nil
}

// Create membuat blog baru dengan file upload.
func (s *BlogService) Create(req dto.CreateBlogRequest, file *multipart.FileHeader) (*models.Blog, error) {
	imagePath, err := utils.SaveUploadedFile(file, "blogs")
	if err != nil {
		return nil, err
	}

	blog := &models.Blog{
		ID:       uuid.New().String(),
		Title:    req.Title,
		Content:  req.Content,
		ImageURL: imagePath,
	}

	if err := s.blogRepo.Create(blog); err != nil {
		utils.DeleteFile(imagePath)
		return nil, errors.New("gagal membuat blog")
	}

	blog.ImageURL = s.buildFileURL(blog.ImageURL)
	return blog, nil
}

// Update memperbarui blog yang sudah ada.
func (s *BlogService) Update(id string, req dto.UpdateBlogRequest, file *multipart.FileHeader) (*models.Blog, error) {
	blog, err := s.blogRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("blog tidak ditemukan")
	}

	oldImagePath := blog.ImageURL

	if req.Title != "" {
		blog.Title = req.Title
	}
	if req.Content != "" {
		blog.Content = req.Content
	}

	if file != nil {
		newImagePath, err := utils.SaveUploadedFile(file, "blogs")
		if err != nil {
			return nil, err
		}
		blog.ImageURL = newImagePath
	}

	if err := s.blogRepo.Update(blog); err != nil {
		if file != nil {
			utils.DeleteFile(blog.ImageURL)
		}
		return nil, errors.New("gagal memperbarui blog")
	}

	if file != nil {
		utils.DeleteFile(oldImagePath)
	}

	blog.ImageURL = s.buildFileURL(blog.ImageURL)
	return blog, nil
}

// Delete menghapus blog berdasarkan ID.
func (s *BlogService) Delete(id string) error {
	blog, err := s.blogRepo.FindByID(id)
	if err != nil {
		return errors.New("blog tidak ditemukan")
	}

	rowsAffected, err := s.blogRepo.Delete(id)
	if err != nil {
		return errors.New("gagal menghapus blog")
	}
	if rowsAffected == 0 {
		return errors.New("blog tidak ditemukan")
	}

	utils.DeleteFile(blog.ImageURL)
	return nil
}

// buildFileURL membangun URL lengkap dari path relatif file.
func (s *BlogService) buildFileURL(relPath string) string {
	if relPath == "" {
		return ""
	}
	return s.baseURL + relPath
}
