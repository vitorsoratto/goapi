package model

import "gorm.io/gorm"

type NotificationApp struct {
	*gorm.Model
	AppName            string
	URL                string
	FirebaseMessageKey *string
	Credentials        *string
}

type NotificationAppResponse struct {
	ID                 uint `json:"id"`
	AppName            string `json:"app_name"`
	URL                string `json:"url"`
	FirebaseMessageKey *string `json:"firebase_message_key"`
	Credentials        *string `json:"credentials"`
}
