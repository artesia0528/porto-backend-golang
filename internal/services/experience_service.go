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

// ExperienceService menangani business logic untuk experience.
type ExperienceService struct {
	experienceRepo *repositories.ExperienceRepository
	baseURL        string
}

// NewExperienceService membuat instance baru ExperienceService.
func NewExperienceService(experienceRepo *repositories.ExperienceRepository, baseURL string) *ExperienceService {
	return &ExperienceService{experienceRepo: experienceRepo, baseURL: baseURL}
}

// GetAll mengambil semua experience.
func (s *ExperienceService) GetAll() ([]models.Experience, error) {
	experiences, err := s.experienceRepo.FindAll()
	if err != nil {
		return nil, err
	}

	for i := range experiences {
		experiences[i].LogoURL = s.buildFileURL(experiences[i].LogoURL)
	}

	return experiences, nil
}

// GetByID mengambil satu experience berdasarkan ID.
func (s *ExperienceService) GetByID(id string) (*models.Experience, error) {
	experience, err := s.experienceRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("experience tidak ditemukan")
	}

	experience.LogoURL = s.buildFileURL(experience.LogoURL)
	return experience, nil
}

// Create membuat experience baru dengan file upload.
func (s *ExperienceService) Create(req dto.CreateExperienceRequest, file *multipart.FileHeader) (*models.Experience, error) {
	logoPath, err := utils.SaveUploadedFile(file, "experiences")
	if err != nil {
		return nil, err
	}

	experience := &models.Experience{
		ID:          uuid.New().String(),
		Company:     req.Company,
		Position:    req.Position,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		IsCurrent:   req.IsCurrent,
		Description: req.Description,
		LogoURL:     logoPath,
	}

	if err := s.experienceRepo.Create(experience); err != nil {
		utils.DeleteFile(logoPath)
		return nil, errors.New("gagal membuat experience")
	}

	experience.LogoURL = s.buildFileURL(experience.LogoURL)
	return experience, nil
}

// Update memperbarui experience yang sudah ada.
func (s *ExperienceService) Update(id string, req dto.UpdateExperienceRequest, file *multipart.FileHeader) (*models.Experience, error) {
	experience, err := s.experienceRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("experience tidak ditemukan")
	}

	oldLogoPath := experience.LogoURL

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

	if file != nil {
		newLogoPath, err := utils.SaveUploadedFile(file, "experiences")
		if err != nil {
			return nil, err
		}
		experience.LogoURL = newLogoPath
	}

	if err := s.experienceRepo.Update(experience); err != nil {
		if file != nil {
			utils.DeleteFile(experience.LogoURL)
		}
		return nil, errors.New("gagal memperbarui experience")
	}

	if file != nil {
		utils.DeleteFile(oldLogoPath)
	}

	experience.LogoURL = s.buildFileURL(experience.LogoURL)
	return experience, nil
}

// Delete menghapus experience berdasarkan ID.
func (s *ExperienceService) Delete(id string) error {
	experience, err := s.experienceRepo.FindByID(id)
	if err != nil {
		return errors.New("experience tidak ditemukan")
	}

	rowsAffected, err := s.experienceRepo.Delete(id)
	if err != nil {
		return errors.New("gagal menghapus experience")
	}
	if rowsAffected == 0 {
		return errors.New("experience tidak ditemukan")
	}

	utils.DeleteFile(experience.LogoURL)
	return nil
}

// buildFileURL membangun URL lengkap dari path relatif file.
func (s *ExperienceService) buildFileURL(relPath string) string {
	if relPath == "" {
		return ""
	}
	return s.baseURL + relPath
}
