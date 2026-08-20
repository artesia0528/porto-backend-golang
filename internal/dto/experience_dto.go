package dto

// CreateExperienceRequest adalah DTO untuk request membuat experience baru (multipart/form-data).
type CreateExperienceRequest struct {
	Company     string `form:"company" binding:"required,min=2,max=100"`
	Position    string `form:"position" binding:"required,min=2,max=100"`
	StartDate   string `form:"start_date" binding:"required"`
	EndDate     string `form:"end_date"`
	IsCurrent   bool   `form:"is_current"`
	Description string `form:"description"`
}

// UpdateExperienceRequest adalah DTO untuk request update experience (multipart/form-data).
type UpdateExperienceRequest struct {
	Company     string `form:"company" binding:"omitempty,min=2,max=100"`
	Position    string `form:"position" binding:"omitempty,min=2,max=100"`
	StartDate   string `form:"start_date"`
	EndDate     string `form:"end_date"`
	IsCurrent   *bool  `form:"is_current"` // pointer agar bisa bedakan false vs tidak diisi
	Description string `form:"description"`
}
