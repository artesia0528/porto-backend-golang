package services

import (
	"errors"
	"portfolio-backend/internal/dto"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"
)

// ProjectService menangani business logic untuk project.
type ProjectService struct {
	projectRepo *repositories.ProjectRepository
}

// NewProjectService membuat instance baru ProjectService.
func NewProjectService(projectRepo *repositories.ProjectRepository) *ProjectService {
	return &ProjectService{projectRepo: projectRepo}
}

// GetAll mengambil semua project.
func (s *ProjectService) GetAll() ([]models.Project, error) {
	return s.projectRepo.FindAll()
}

// GetByID mengambil satu project berdasarkan ID.
func (s *ProjectService) GetByID(id uint) (*models.Project, error) {
	project, err := s.projectRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("project tidak ditemukan")
	}
	return project, nil
}

// Create membuat project baru.
func (s *ProjectService) Create(req dto.CreateProjectRequest) (*models.Project, error) {
	project := &models.Project{
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    req.ImageURL,
	}

	if err := s.projectRepo.Create(project); err != nil {
		return nil, errors.New("gagal membuat project")
	}

	return project, nil
}

// Update memperbarui project yang sudah ada.
func (s *ProjectService) Update(id uint, req dto.UpdateProjectRequest) (*models.Project, error) {
	project, err := s.projectRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("project tidak ditemukan")
	}

	// Update hanya field yang diisi
	if req.Title != "" {
		project.Title = req.Title
	}
	if req.Description != "" {
		project.Description = req.Description
	}
	if req.ImageURL != "" {
		project.ImageURL = req.ImageURL
	}

	if err := s.projectRepo.Update(project); err != nil {
		return nil, errors.New("gagal memperbarui project")
	}

	return project, nil
}

// Delete menghapus project berdasarkan ID.
func (s *ProjectService) Delete(id uint) error {
	rowsAffected, err := s.projectRepo.Delete(id)
	if err != nil {
		return errors.New("gagal menghapus project")
	}
	if rowsAffected == 0 {
		return errors.New("project tidak ditemukan")
	}
	return nil
}
