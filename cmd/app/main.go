package main

import (
	"calpal-core/database"
	"calpal-core/handler/auth"
	"calpal-core/handler/user"
	"calpal-core/repository"
	"fmt"
	"log"
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

	authRepository := repository.NewAuthRepository(db)
	authHandler := auth.NewAuthHandler(authRepository)

	userRepository := repository.NewUserRepository(db)
	userHandler := user.New(userRepository)

	http.HandleFunc("/sign-up", authHandler.SignUp)
	http.HandleFunc("/sign-in", authHandler.SignIn)

	http.HandleFunc("/users", userHandler.Users)

	// Set target calories
	http.HandleFunc("/set-target-calories", setTargetCalories)

	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
	// food log
}
