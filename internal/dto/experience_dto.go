package dto

// CreateExperienceRequest adalah DTO untuk request membuat experience baru.
type CreateExperienceRequest struct {
	Company     string `json:"company" binding:"required,min=2,max=100"`
	Position    string `json:"position" binding:"required,min=2,max=100"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date"`
	IsCurrent   bool   `json:"is_current"`
	Description string `json:"description"`
	LogoURL     string `json:"logo_url" binding:"omitempty,url"`
}

// UpdateExperienceRequest adalah DTO untuk request update experience.
type UpdateExperienceRequest struct {
	Company     string `json:"company" binding:"omitempty,min=2,max=100"`
	Position    string `json:"position" binding:"omitempty,min=2,max=100"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	IsCurrent   *bool  `json:"is_current"` // pointer agar bisa bedakan false vs tidak diisi
	Description string `json:"description"`
	LogoURL     string `json:"logo_url" binding:"omitempty,url"`
}
