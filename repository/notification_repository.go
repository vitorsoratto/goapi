package repository

import (
	"errors"

	"goapi/model"

	"gorm.io/gorm"
)

type NotificationRepository struct {
	connection *gorm.DB
}

func NewNotificationRepository(connection *gorm.DB) *NotificationRepository {
	return &NotificationRepository{
		connection: connection,
	}
}

func (r *NotificationRepository) GetList() ([]model.Notification, error) {
	notifications := []model.Notification{}
	if err := r.connection.Find(&notifications).Error; err != nil {
		return nil, errors.New("Error while listing notifications on the database")
	}

	return notifications, nil
}
