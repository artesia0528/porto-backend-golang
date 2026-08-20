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

// ProjectService menangani business logic untuk project.
type ProjectService struct {
	projectRepo *repositories.ProjectRepository
	baseURL     string
}

// NewProjectService membuat instance baru ProjectService.
func NewProjectService(projectRepo *repositories.ProjectRepository, baseURL string) *ProjectService {
	return &ProjectService{projectRepo: projectRepo, baseURL: baseURL}
}

// GetAll mengambil semua project.
func (s *ProjectService) GetAll() ([]models.Project, error) {
	projects, err := s.projectRepo.FindAll()
	if err != nil {
		return nil, err
	}

	// Bangun full URL untuk setiap project
	for i := range projects {
		projects[i].ImageURL = s.buildFileURL(projects[i].ImageURL)
	}

	return projects, nil
}

// GetByID mengambil satu project berdasarkan ID.
func (s *ProjectService) GetByID(id string) (*models.Project, error) {
	project, err := s.projectRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("project tidak ditemukan")
	}

	project.ImageURL = s.buildFileURL(project.ImageURL)
	return project, nil
}

// Create membuat project baru dengan file upload.
func (s *ProjectService) Create(req dto.CreateProjectRequest, file *multipart.FileHeader) (*models.Project, error) {
	imagePath, err := utils.SaveUploadedFile(file, "projects")
	if err != nil {
		return nil, err
	}

	project := &models.Project{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    imagePath,
	}

	if err := s.projectRepo.Create(project); err != nil {
		utils.DeleteFile(imagePath) // Bersihkan file jika gagal simpan ke DB
		return nil, errors.New("gagal membuat project")
	}

	project.ImageURL = s.buildFileURL(project.ImageURL)
	return project, nil
}

// Update memperbarui project yang sudah ada.
func (s *ProjectService) Update(id string, req dto.UpdateProjectRequest, file *multipart.FileHeader) (*models.Project, error) {
	project, err := s.projectRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("project tidak ditemukan")
	}

	// Simpan path lama untuk dihapus nanti
	oldImagePath := project.ImageURL

	// Update field yang diisi
	if req.Title != "" {
		project.Title = req.Title
	}
	if req.Description != "" {
		project.Description = req.Description
	}

	// Upload file baru jika ada
	if file != nil {
		newImagePath, err := utils.SaveUploadedFile(file, "projects")
		if err != nil {
			return nil, err
		}
		project.ImageURL = newImagePath
	}

	if err := s.projectRepo.Update(project); err != nil {
		// Jika file baru diupload tapi gagal simpan, hapus file baru
		if file != nil {
			utils.DeleteFile(project.ImageURL)
		}
		return nil, errors.New("gagal memperbarui project")
	}

	// Hapus file lama setelah update berhasil
	if file != nil {
		utils.DeleteFile(oldImagePath)
	}

	project.ImageURL = s.buildFileURL(project.ImageURL)
	return project, nil
}

// Delete menghapus project berdasarkan ID.
func (s *ProjectService) Delete(id string) error {
	project, err := s.projectRepo.FindByID(id)
	if err != nil {
		return errors.New("project tidak ditemukan")
	}

	rowsAffected, err := s.projectRepo.Delete(id)
	if err != nil {
		return errors.New("gagal menghapus project")
	}
	if rowsAffected == 0 {
		return errors.New("project tidak ditemukan")
	}

	// Hapus file dari disk
	utils.DeleteFile(project.ImageURL)

	return nil
}

// buildFileURL membangun URL lengkap dari path relatif file.
func (s *ProjectService) buildFileURL(relPath string) string {
	if relPath == "" {
		return ""
	}
	return s.baseURL + relPath
}
