package handlers

import (
	"net/http"
	"portfolio-backend/internal/dto"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// AuthHandler menangani HTTP request untuk autentikasi.
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler membuat instance baru AuthHandler.
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login menangani POST /api/login.
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Input tidak valid: "+err.Error())
		return
	}

	token, err := h.authService.Login(req)
	if err != nil {
		models.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	models.SuccessResponse(c, http.StatusOK, "Login berhasil", gin.H{"token": token})
}

// Register menangani POST /api/register.
// func (h *AuthHandler) Register(c *gin.Context) {
// 	var req dto.RegisterRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		models.ErrorResponse(c, http.StatusBadRequest, "Input tidak valid: "+err.Error())
// 		return
// 	}

// 	user, err := h.authService.Register(req)
// 	if err != nil {
// 		models.ErrorResponse(c, http.StatusConflict, err.Error())
// 		return
// 	}

// 	models.SuccessResponse(c, http.StatusCreated, "User berhasil dibuat", gin.H{
// 		"id":       user.ID,
// 		"username": user.Username,
// 	})
// }