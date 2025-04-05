package main

import (
	"calpal-core/database"
	"calpal-core/handler"
	"calpal-core/repository"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func setTargetCalories(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented")
}

func main() {
	// Connect to database
	db := database.ConnectToPostgres()

	defer database.Close(db)

	// Execute migrations
	database.ExecuteMigrations(db)

	authRepository := repository.NewAuthRepository(db)

	authHandler := handler.NewAuthHandler(authRepository)

	http.HandleFunc("/sign-up", authHandler.SignUp)
	http.HandleFunc("/sign-in", authHandler.SignIn)

	// Set target calories
	http.HandleFunc("/set-target-calories", setTargetCalories)

	http.ListenAndServe(":8080", nil)
	// food log
}
