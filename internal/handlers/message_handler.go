package handlers

import (
	"net/http"
	"portfolio-backend/internal/dto"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	messageService *services.MessageService
}

func NewMessageHandler(messageService *services.MessageService) *MessageHandler {
	return &MessageHandler{messageService: messageService}
}

// Create menangani POST /api/contact (publik)
func (h *MessageHandler) Create(c *gin.Context) {
	var req dto.CreateMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Input tidak valid: "+err.Error())
		return
	}

	message, err := h.messageService.Create(req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengirim pesan")
		return
	}

	models.SuccessResponse(c, http.StatusCreated, "Pesan berhasil dikirim", message)
}

// GetAll menangani GET /api/admin/messages (butuh login)
func (h *MessageHandler) GetAll(c *gin.Context) {
	messages, err := h.messageService.GetAll()
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil pesan")
		return
	}
	models.SuccessResponse(c, http.StatusOK, "Berhasil", messages)
}

// MarkAsRead menangani PATCH /api/admin/messages/:id/read
func (h *MessageHandler) MarkAsRead(c *gin.Context) {
	id := c.Param("id")
	if err := h.messageService.MarkAsRead(id); err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal update pesan")
		return
	}
	models.SuccessResponse(c, http.StatusOK, "Pesan ditandai sudah dibaca", nil)
}

// Delete menangani DELETE /api/admin/messages/:id
func (h *MessageHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.messageService.Delete(id); err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal menghapus pesan")
		return
	}
	models.SuccessResponse(c, http.StatusOK, "Pesan dihapus", nil)
}