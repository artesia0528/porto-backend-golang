package handlers

import (
	"net/http"
	"portfolio-backend/internal/dto"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	blogService *services.BlogService
}

func NewBlogHandler(blogService *services.BlogService) *BlogHandler {
	return &BlogHandler{blogService: blogService}
}

// GET /api/blogs (publik)
func (h *BlogHandler) GetAll(c *gin.Context) {
	blogs, err := h.blogService.GetAll()
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil data blog")
		return
	}
	models.SuccessResponse(c, http.StatusOK, "Berhasil", blogs)
}

// POST /api/admin/blogs (butuh login)
func (h *BlogHandler) Create(c *gin.Context) {
	var req dto.CreateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Input tidak valid: "+err.Error())
		return
	}

	blog, err := h.blogService.Create(req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	models.SuccessResponse(c, http.StatusCreated, "Blog berhasil dibuat", blog)
}

// PUT /api/admin/blogs/:id (butuh login)
func (h *BlogHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Input tidak valid: "+err.Error())
		return
	}

	blog, err := h.blogService.Update(id, req)
	if err != nil {
		models.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	models.SuccessResponse(c, http.StatusOK, "Blog berhasil diperbarui", blog)
}

// DELETE /api/admin/blogs/:id (butuh login)
func (h *BlogHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.blogService.Delete(id); err != nil {
		models.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	models.SuccessResponse(c, http.StatusOK, "Blog berhasil dihapus", nil)
}