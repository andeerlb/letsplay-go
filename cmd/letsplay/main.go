package main

import (
	"letsplay-microservice/internal/bootstrap"
	"letsplay-microservice/internal/config"
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

	container := bootstrap.BuildContainer(cfg, logg)
	router := router.NewRouter(container, logg, cfg)

	server := server.NewServer(router, cfg.ServerPort, logg)
	if err := server.Start(); err != nil {
		logg.Fatal("server failed to start", zap.Error(err))
	}
}
