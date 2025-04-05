package database

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "testuser"
	password = "testpassword"
	dbname   = "calpal-db"
)

func ConnectToPostgres() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to database")

	return db
}

func ExecuteMigrations(db *sql.DB) {
	log.Println("Executing migrations")

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (" +
		"id VARCHAR(36) PRIMARY KEY," +
		"first_name VARCHAR(100) NOT NULL," +
		"last_name VARCHAR(100) NOT NULL," +
		"password VARCHAR(255) NOT NULL," +
		"email VARCHAR(255) UNIQUE NOT NULL," +
		"target_calories INT," +
		"created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	log.Println("Successfully executed migrations")
}

func Close(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
