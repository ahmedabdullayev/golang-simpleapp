package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // postgres golang driver
	"log"
	"os"
	"strconv"
)

var db *sql.DB

func dsn() string {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}

func Init() (*sql.DB, error) {
	connection, err := sql.Open("postgres", dsn())
	if err != nil {
		return connection, err
	}

	err = connection.Ping()
	if err != nil {
		return connection, err
	}

	log.Println("Database connection established")

	return connection, nil
}
