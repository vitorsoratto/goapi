package model

import "gorm.io/gorm"

type NotificationApp struct {
	*gorm.Model
	AppName            string
	URL                string
	FirebaseMessageKey *string
	Credentials        *string
}
