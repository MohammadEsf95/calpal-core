package main

import (
	"calpal-core/database"
	"calpal-core/pkg/config_loader"
	httpserver "calpal-core/pkg/http_server"
	"calpal-core/pkg/postgresmigrator"
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
	PostgresDB database.Config   `koanf:"postgres"`
	HttpServer httpserver.Config `koanf:"http_server"`
}

func main() {

	var config Config
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error loading working directory: %v", err)
	}

	err = configloader.Load(filepath.Join(workDir, "deploy", "config.yml"), &config)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Connect to database
	db := database.ConnectToPostgres(config.PostgresDB)

	defer database.Close(db)

	// Execute the migrations
	migrator := postgresmigrator.New(config.PostgresDB, filepath.Join(workDir, "database", "migration"))
	migrator.Up()

	app := Setup(db, config)
	app.Start()
}
