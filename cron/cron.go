package cron

import (
	"fmt"
	"time"

	"goapi/config"
	"goapi/model"
	"goapi/repository"

	"github.com/go-co-op/gocron/v2"
)

func getNotifications() []model.NotificationApp {
	notificationAppRepository := repository.NewNotificationAppRepository(config.Database)
	notificationApps := []model.NotificationApp{}
	notificationApps, err := notificationAppRepository.GetList()
	if err != nil {
		fmt.Printf("Error getting the notification apps: %v", err)
		return nil
	}

	return notificationApps
}

func InitCron() {
	s, err := gocron.NewScheduler()
	if err != nil {
		fmt.Printf("Error starting the scheduler: %v", err)
	}

	_, err = s.NewJob(
		gocron.DurationJob(time.Minute),
		gocron.NewTask(
			func(notificationApps []model.NotificationApp) {
				fmt.Println("--------- NOTIFICATIONS ---------")
				fmt.Printf("%+v\n", notificationApps)
				fmt.Println("---------------------------------")
			},
			getNotifications(),
		),
	)
	if err != nil {
		fmt.Printf("Error creating the job: %v", err)
	}

	s.Start()
}
