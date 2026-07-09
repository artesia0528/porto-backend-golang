package models

import "github.com/gin-gonic/gin"

// APIResponse adalah wrapper standar untuk semua API response.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse mengirim response sukses dengan format standar.
func SuccessResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse mengirim response error dengan format standar.
func ErrorResponse(c *gin.Context, status int, errMsg string) {
	c.JSON(status, APIResponse{
		Success: false,
		Error:   errMsg,
	})
}
