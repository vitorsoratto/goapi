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

type NotificationResponse struct {
	ID            int        `json:"id"`
	AppID         int        `json:"app_id"`
	Title         string     `json:"title"`
	Body          string     `json:"body"`
	Image         *string    `json:"image"`
	ToChannel     string     `json:"to_channel"`
	LastExecution *time.Time `json:"last_execution"`
	ExecuteAt     *string    `json:"execute_at"`
	Sun           string     `json:"sun"`
	Mon           string     `json:"mon"`
	Tue           string     `json:"tue"`
	Wed           string     `json:"wed"`
	Thu           string     `json:"thu"`
	Fri           string     `json:"fri"`
	Sat           string     `json:"sat"`
	Active        bool       `json:"active"`
	MonthDay      *uint8     `json:"monthDay"`
	URLLauncher   *string    `json:"url_launcher"`
	TimesSent     uint       `json:"times_sent"`
}
