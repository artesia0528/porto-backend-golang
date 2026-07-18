package services

import (
	"errors"
	"log"
	"portfolio-backend/internal/dto"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"

	"github.com/google/uuid"
)

type MessageService struct {
	messageRepo *repositories.MessageRepository
}

func NewMessageService(messageRepo *repositories.MessageRepository) *MessageService {
	return &MessageService{messageRepo: messageRepo}
}

func (s *MessageService) Create(req dto.CreateMessageRequest) (*models.Message, error) {
	message := &models.Message{
		ID:      uuid.New().String(),
		Name:    req.Name,
		Email:   req.Email,
		Subject: req.Subject,
		Content: req.Content,
	}

	if err := s.messageRepo.Create(message); err != nil {
		return nil, errors.New("gagal menyimpan pesan")
	}

	// Kirim notifikasi di background (goroutine) — tidak blocking response ke user
	go s.sendNotification(*message)

	return message, nil
}

func (s *MessageService) GetAll() ([]models.Message, error) {
	return s.messageRepo.FindAll()
}

func (s *MessageService) MarkAsRead(id string) error {
	return s.messageRepo.UpdateReadStatus(id, true)
}

func (s *MessageService) Delete(id string) error {
	rowsAffected, err := s.messageRepo.Delete(id)
	if err != nil {
		return errors.New("gagal menghapus pesan")
	}
	if rowsAffected == 0 {
		return errors.New("pesan tidak ditemukan")
	}
	return nil
}

// sendNotification simulasi kirim email — nanti bisa diganti pakai SMTP/service pihak ketiga
func (s *MessageService) sendNotification(message models.Message) {
	log.Printf("Notifikasi: pesan baru dari %s (%s)\n", message.Name, message.Email)
	// TODO: integrasi SMTP/Telegram bot di sini
}