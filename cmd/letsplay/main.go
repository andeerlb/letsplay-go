package main

import (
	"letsplay-microservice/internal/bootstrap"
	"letsplay-microservice/internal/config"
	"letsplay-microservice/internal/database"
	"letsplay-microservice/internal/logger"
	"letsplay-microservice/internal/router"
	"letsplay-microservice/internal/server"
	"log"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logg, err := logger.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	database.InitPostgres(cfg)

	db, err := database.InitPostgres(cfg)
	if err != nil {
		logg.Fatal("failed to connect to postgres", zap.Error(err))
	}
	defer db.Close()

	database.RunMigrations(db.DB)

	container := bootstrap.BuildContainer(cfg, logg)
	r := router.NewRouter(container, logg, cfg)

	s := server.NewServer(r, cfg.ServerPort, logg)
	if err := s.Start(); err != nil {
		logg.Fatal("server failed to start", zap.Error(err))
	}
}
