package handlers

import (
	"net/http"
	"portfolio-backend/internal/dto"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// ExperienceHandler menangani HTTP request untuk experience.
type ExperienceHandler struct {
	experienceService *services.ExperienceService
}

// NewExperienceHandler membuat instance baru ExperienceHandler.
func NewExperienceHandler(experienceService *services.ExperienceService) *ExperienceHandler {
	return &ExperienceHandler{experienceService: experienceService}
}

// GetAll menangani GET /api/experiences (publik).
func (h *ExperienceHandler) GetAll(c *gin.Context) {
	experiences, err := h.experienceService.GetAll()
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil data experience")
		return
	}

	models.SuccessResponse(c, http.StatusOK, "Berhasil", experiences)
}

// Create menangani POST /api/admin/experiences (butuh login).
func (h *ExperienceHandler) Create(c *gin.Context) {
	var req dto.CreateExperienceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Input tidak valid: "+err.Error())
		return
	}

	experience, err := h.experienceService.Create(req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.SuccessResponse(c, http.StatusCreated, "Experience berhasil dibuat", experience)
}

// Update menangani PUT /api/admin/experiences/:id (butuh login).
func (h *ExperienceHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		models.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req dto.UpdateExperienceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Input tidak valid: "+err.Error())
		return
	}

	experience, err := h.experienceService.Update(id, req)
	if err != nil {
		models.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	models.SuccessResponse(c, http.StatusOK, "Experience berhasil diperbarui", experience)
}

// Delete menangani DELETE /api/admin/experiences/:id (butuh login).
func (h *ExperienceHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		models.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := h.experienceService.Delete(id); err != nil {
		models.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	models.SuccessResponse(c, http.StatusOK, "Experience berhasil dihapus", nil)
}
