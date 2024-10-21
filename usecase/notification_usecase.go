package usecase

import (
	"goapi/model"
	"goapi/repository"
)

type NotificationUsecase struct {
	repository *repository.NotificationRepository
}

func NewNotificationUsecase(repository *repository.NotificationRepository) *NotificationUsecase {
	return &NotificationUsecase{
		repository: repository,
	}
}

func (u *NotificationUsecase) GetList() ([]model.Notification, error) {
	notifications, err := u.repository.GetList()
	if err != nil {
		return nil, err
	}

	return notifications, nil
}
