package handlers

import (
	"net/http"
	"portfolio-backend/internal/dto"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// ProjectHandler menangani HTTP request untuk project.
type ProjectHandler struct {
	projectService *services.ProjectService
}

// NewProjectHandler membuat instance baru ProjectHandler.
func NewProjectHandler(projectService *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{projectService: projectService}
}

// GetProjects menangani GET /api/projects (publik).
func (h *ProjectHandler) GetProjects(c *gin.Context) {
	projects, err := h.projectService.GetAll()
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil data project")
		return
	}

	models.SuccessResponse(c, http.StatusOK, "Berhasil", projects)
}

// CreateProject menangani POST /api/admin/projects (butuh login).
// Menerima multipart/form-data: title, description, image (file).
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var req dto.CreateProjectRequest
	if err := c.ShouldBind(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Input tidak valid: "+err.Error())
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "File gambar wajib diunggah")
		return
	}

	project, err := h.projectService.Create(req, file)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.SuccessResponse(c, http.StatusCreated, "Project berhasil dibuat", project)
}

// UpdateProject menangani PUT /api/admin/projects/:id (butuh login).
// Menerima multipart/form-data: title, description, image (file, optional).
func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		models.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req dto.UpdateProjectRequest
	if err := c.ShouldBind(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Input tidak valid: "+err.Error())
		return
	}

	// File optional saat update
	file, _ := c.FormFile("image")

	project, err := h.projectService.Update(id, req, file)
	if err != nil {
		models.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	models.SuccessResponse(c, http.StatusOK, "Project berhasil diperbarui", project)
}

// DeleteProject menangani DELETE /api/admin/projects/:id (butuh login).
func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		models.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := h.projectService.Delete(id); err != nil {
		models.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	models.SuccessResponse(c, http.StatusOK, "Project berhasil dihapus", nil)
}
