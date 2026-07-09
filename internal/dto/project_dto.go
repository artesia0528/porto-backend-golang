package dto

// CreateProjectRequest adalah DTO untuk request membuat project baru.
type CreateProjectRequest struct {
	Title       string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=10"`
	ImageURL    string `json:"image_url" binding:"omitempty,url"`
}

// UpdateProjectRequest adalah DTO untuk request update project.
type UpdateProjectRequest struct {
	Title       string `json:"title" binding:"omitempty,min=3,max=100"`
	Description string `json:"description" binding:"omitempty,min=10"`
	ImageURL    string `json:"image_url" binding:"omitempty,url"`
}
