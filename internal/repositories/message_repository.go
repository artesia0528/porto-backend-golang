package repositories

import (
	"portfolio-backend/internal/models"

	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(message *models.Message) error {
	return r.db.Create(message).Error
}

func (r *MessageRepository) FindAll() ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Order("created_at DESC").Find(&messages).Error
	return messages, err
}

func (r *MessageRepository) UpdateReadStatus(id string, isRead bool) error {
	return r.db.Model(&models.Message{}).Where("id = ?", id).Update("is_read", isRead).Error
}

func (r *MessageRepository) Delete(id string) (int64, error) {
	result := r.db.Where("id = ?", id).Delete(&models.Message{})
	return result.RowsAffected, result.Error
}
