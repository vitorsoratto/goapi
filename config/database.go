package config

import (
	"fmt"
	"os"

	"goapi/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDB() error {
	host := os.Getenv("GOAPI_DB_HOST")
	port := os.Getenv("GOAPI_DB_PORT")
	user := os.Getenv("GOAPI_DB_USER")
	pass := os.Getenv("GOAPI_DB_PASS")
	dbname := os.Getenv("GOAPI_DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=America/Sao_Paulo",
		host, user, pass, dbname, port)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  psqlInfo,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	Database = db

	fmt.Println("Connected to PostgreSQL on " + dbname)

	err = db.AutoMigrate(&model.NotificationApp{}, &model.Notification{})

	return nil
}
