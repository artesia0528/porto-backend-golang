package dto

type CreateBlogRequest struct {
	Title    string `json:"title" binding:"required,min=3,max=100"`
	Content  string `json:"content" binding:"required,min=10"`
	ImageURL string `json:"image_url" binding:"omitempty,url"`
}

type UpdateBlogRequest struct {
	Title    string `json:"title" binding:"omitempty,min=3,max=100"`
	Content  string `json:"content" binding:"omitempty,min=10"`
	ImageURL string `json:"image_url" binding:"omitempty,url"`
}