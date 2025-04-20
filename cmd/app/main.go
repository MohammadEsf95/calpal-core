package main

import (
	"calpal-core/database"
	"calpal-core/handler/auth"
	"calpal-core/handler/user"
	"calpal-core/pkg/config_loader"
	"calpal-core/repository"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

func setTargetCalories(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented")
}

type Config struct {
	PostgresDB database.Config `koanf:"postgres"`
}

func main() {

	var config Config
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error loading working directory: %v", err)
	}

	err = configloader.Load(filepath.Join(workDir, "database", "dbconfig.yml"), &config)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	fmt.Println(config.PostgresDB)

	// Connect to database
	db := database.ConnectToPostgres(config.PostgresDB)

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
