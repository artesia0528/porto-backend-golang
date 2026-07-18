package dto

type CreateMessageRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Subject string `json:"subject" binding:"required"`
	Content string `json:"content" binding:"required,min=10"`
}