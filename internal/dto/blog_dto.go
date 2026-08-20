package dto

// CreateBlogRequest adalah DTO untuk request membuat blog baru (multipart/form-data).
type CreateBlogRequest struct {
	Title   string `form:"title" binding:"required,min=3,max=100"`
	Content string `form:"content" binding:"required,min=10"`
}

// UpdateBlogRequest adalah DTO untuk request update blog (multipart/form-data).
type UpdateBlogRequest struct {
	Title   string `form:"title" binding:"omitempty,min=3,max=100"`
	Content string `form:"content" binding:"omitempty,min=10"`
}
