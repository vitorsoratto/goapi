package controller

import (
	"net/http"

	"goapi/usecase"

	"github.com/gin-gonic/gin"
)

type NotificationAppController struct {
	useCase *usecase.NotificationAppUsecase
}

func NewNotificationAppController(usecase *usecase.NotificationAppUsecase) *NotificationAppController {
	return &NotificationAppController{
		useCase: usecase,
	}
}

func (c *NotificationAppController) GetList(ctx *gin.Context) {
	notificationApps, err := c.useCase.GetList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, notificationApps)
}
