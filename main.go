package main

import (
	"calpal-core/entity"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "testuser"
	password = "testpassword"
	dbname   = "calpal-db"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user entity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{Message: err.Error()}
		bytes, _ := json.Marshal(errorResponse)
		w.Write(bytes)
		return
	}

	fmt.Printf("User %+v\n", user)
	// save to db

	json.NewEncoder(w).Encode(user)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sign In")
}

func setTargetCalories(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented")
}

func connectToPostgres() *sql.DB {
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
func main() {
	// Connect to database
	db := connectToPostgres()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	// User registration: signup / login
	http.HandleFunc("/sign-up", signUp)
	http.HandleFunc("/sign-in", signIn)

	// Set target calories
	http.HandleFunc("/set-target-calories", setTargetCalories)

	http.ListenAndServe(":8080", nil)
	// food log
}
