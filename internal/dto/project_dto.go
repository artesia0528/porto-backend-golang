package dto

// CreateProjectRequest adalah DTO untuk request membuat project baru (multipart/form-data).
type CreateProjectRequest struct {
	Title       string `form:"title" binding:"required,min=3,max=100"`
	Description string `form:"description" binding:"required,min=10"`
}

// UpdateProjectRequest adalah DTO untuk request update project (multipart/form-data).
type UpdateProjectRequest struct {
	Title       string `form:"title" binding:"omitempty,min=3,max=100"`
	Description string `form:"description" binding:"omitempty,min=10"`
}
