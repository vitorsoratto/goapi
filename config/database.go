package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var Database *sql.DB

func ConnectDB() error {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	dbname := os.Getenv("DATABASE_NAME")

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, pass, host, port, dbname)

	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		return err
	}

	Database = db

	fmt.Println("Connected to PostgreSQL on " + dbname)

	return nil
}
