package model

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	*gorm.Model
	AppID         int
	App           NotificationApp
	Title         string
	Body          string
	Image         *string
	ToChannel     string
	LastExecution *time.Time
	ExecuteAt     *string
	Sun           string `gorm:"size:1,default:N"`
	Mon           string `gorm:"size:1,default:N"`
	Tue           string `gorm:"size:1,default:N"`
	Wed           string `gorm:"size:1,default:N"`
	Thu           string `gorm:"size:1,default:N"`
	Fri           string `gorm:"size:1,default:N"`
	Sat           string `gorm:"size:1,default:N"`
	Active        bool   `gorm:"default:false"`
	MonthDay      *uint8
	URLLauncher   *string
	TimesSent     uint `gorm:"default:0"`
}
