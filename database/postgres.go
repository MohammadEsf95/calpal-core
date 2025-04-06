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

func Close(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
