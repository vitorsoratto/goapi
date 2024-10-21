package usecase

import (
	"goapi/model"
	"goapi/repository"
)

type NotificationAppUsecase struct {
	repository *repository.NotificationAppRepository
}

func NewNotificationAppUsecase(repository *repository.NotificationAppRepository) *NotificationAppUsecase {
	return &NotificationAppUsecase{
		repository: repository,
	}
}

func (u *NotificationAppUsecase) GetList() ([]model.NotificationApp, error) {
	notificationApps, err := u.repository.GetList()
	if err != nil {
		return nil, err
	}

	return notificationApps, nil
}
