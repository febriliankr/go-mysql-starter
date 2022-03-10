package helper

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func InitializeDB() (*sqlx.DB, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	var MYSQL_USER,
		MYSQL_PASSWORD,
		MYSQL_HOST,
		MYSQL_PORT, PRODUCTION string

	if PRODUCTION == "true" {
		MYSQL_USER = os.Getenv("MYSQL_USER")
		MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
		MYSQL_HOST = os.Getenv("MYSQL_HOST")
		MYSQL_PORT = os.Getenv("MYSQL_PORT")
	} else {
		MYSQL_USER = os.Getenv("MYSQL_USER_DEV")
		MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD_DEV")
		MYSQL_HOST = os.Getenv("MYSQL_HOST_DEV")
		MYSQL_PORT = os.Getenv("MYSQL_PORT_DEV")
	}

	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")

	connString := fmt.Sprintf("%s:%s@(%s:%s)/%s?tls=true", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE)

	db, err := sqlx.Connect("mysql", connString)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}
