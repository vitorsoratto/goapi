package repository

import (
	"errors"

	"goapi/model"

	"gorm.io/gorm"
)

type NotificationAppRepository struct {
	connection *gorm.DB
}

func NewNotificationAppRepository(connection *gorm.DB) *NotificationAppRepository {
	return &NotificationAppRepository{
		connection: connection,
	}
}

func (r *NotificationAppRepository) GetList() ([]model.NotificationApp, error) {
	notificationApps := []model.NotificationApp{}
	if err := r.connection.Find(&notificationApps).Error; err != nil {
		return nil, errors.New("Error while listing notification apps on the database")
	}

	return notificationApps, nil
}
