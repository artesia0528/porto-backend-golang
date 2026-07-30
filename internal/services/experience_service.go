package services

import (
	"errors"
	"portfolio-backend/internal/dto"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"

	"github.com/google/uuid"
)

// ExperienceService menangani business logic untuk experience.
type ExperienceService struct {
	experienceRepo *repositories.ExperienceRepository
}

// NewExperienceService membuat instance baru ExperienceService.
func NewExperienceService(experienceRepo *repositories.ExperienceRepository) *ExperienceService {
	return &ExperienceService{experienceRepo: experienceRepo}
}

// GetAll mengambil semua experience.
func (s *ExperienceService) GetAll() ([]models.Experience, error) {
	return s.experienceRepo.FindAll()
}

// GetByID mengambil satu experience berdasarkan ID.
func (s *ExperienceService) GetByID(id string) (*models.Experience, error) {
	experience, err := s.experienceRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("experience tidak ditemukan")
	}
	return experience, nil
}

// Create membuat experience baru.
func (s *ExperienceService) Create(req dto.CreateExperienceRequest) (*models.Experience, error) {
	experience := &models.Experience{
		ID:          uuid.New().String(),
		Company:     req.Company,
		Position:    req.Position,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		IsCurrent:   req.IsCurrent,
		Description: req.Description,
		LogoURL:     req.LogoURL,
	}

	if err := s.experienceRepo.Create(experience); err != nil {
		return nil, errors.New("gagal membuat experience")
	}

	return experience, nil
}

// Update memperbarui experience yang sudah ada.
func (s *ExperienceService) Update(id string, req dto.UpdateExperienceRequest) (*models.Experience, error) {
	experience, err := s.experienceRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("experience tidak ditemukan")
	}

	// Update hanya field yang diisi
	if req.Company != "" {
		experience.Company = req.Company
	}
	if req.Position != "" {
		experience.Position = req.Position
	}
	if req.StartDate != "" {
		experience.StartDate = req.StartDate
	}
	if req.EndDate != "" {
		experience.EndDate = req.EndDate
	}
	if req.IsCurrent != nil {
		experience.IsCurrent = *req.IsCurrent
	}
	if req.Description != "" {
		experience.Description = req.Description
	}
	if req.LogoURL != "" {
		experience.LogoURL = req.LogoURL
	}

	if err := s.experienceRepo.Update(experience); err != nil {
		return nil, errors.New("gagal memperbarui experience")
	}

	return experience, nil
}

// Delete menghapus experience berdasarkan ID.
func (s *ExperienceService) Delete(id string) error {
	rowsAffected, err := s.experienceRepo.Delete(id)
	if err != nil {
		return errors.New("gagal menghapus experience")
	}
	if rowsAffected == 0 {
		return errors.New("experience tidak ditemukan")
	}
	return nil
}
