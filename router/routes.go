package router

import (
	"goapi/config"
	"goapi/controller"
	"goapi/repository"
	"goapi/usecase"

	"github.com/gin-gonic/gin"
)

const basePath = "/api"

func initRoutes(router *gin.Engine) {
	initStructure()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Healthy",
		})
	})

	api := router.Group(basePath)
	{
		// Notification Apps
		api.GET("/notification-apps", notificationAppController.GetList)

		// Notification
		api.GET("/notifications", notificationController.GetList)
	}
}

var (
	notificationAppRepository *repository.NotificationAppRepository
	notificationAppUsecase *usecase.NotificationAppUsecase
	notificationAppController *controller.NotificationAppController


	notificationRepository *repository.NotificationRepository
	notificationUsecase *usecase.NotificationUsecase
	notificationController *controller.NotificationController
)

func initStructure() {
	// Repositories
	notificationAppRepository = repository.NewNotificationAppRepository(config.Database)
	notificationRepository = repository.NewNotificationRepository(config.Database)

	// Usecases
	notificationAppUsecase = usecase.NewNotificationAppUsecase(notificationAppRepository)
	notificationUsecase = usecase.NewNotificationUsecase(notificationRepository)

	// Controllers
	notificationAppController = controller.NewNotificationAppController(notificationAppUsecase)
	notificationController = controller.NewNotificationController(notificationUsecase)
}
