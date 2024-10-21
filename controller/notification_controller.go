package controller

import (
	"net/http"

	"goapi/usecase"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	useCase *usecase.NotificationUsecase
}

func NewNotificationController(usecase *usecase.NotificationUsecase) *NotificationController {
	return &NotificationController{
		useCase: usecase,
	}
}

func (c *NotificationController) GetList(ctx *gin.Context) {
	notifications, err := c.useCase.GetList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, notifications)
}
