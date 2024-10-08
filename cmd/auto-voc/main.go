package main

import (
	"log"

	"github.com/Benjosh95/auto-voc/internal/api"
	"github.com/Benjosh95/auto-voc/internal/config"
	"github.com/Benjosh95/auto-voc/internal/db"
	"github.com/Benjosh95/auto-voc/internal/server"
	"github.com/Benjosh95/auto-voc/internal/services"
	"github.com/go-playground/validator/v10"
)

func main() {

	// Load Config
	cfg := config.LoadConfig()

	// Init DB connection
	dbConn := db.InitDB(cfg.DBConfig.ConnectionString)
	defer dbConn.Close()

	// Init dependencies that are needed in the services for example?
	// - like messagebrokers

	// Init services (with dependencies from above like messagebrokers)
	vocService := services.NewVocService(dbConn)

	// Init validator
	validate := validator.New()

	// Init API router (add all services)
	router := api.NewRouter(vocService, validate)

	// Init and run Server
	server := server.NewServer(cfg.ServerConfig, router)
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
