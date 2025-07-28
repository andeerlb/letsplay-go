package bootstrap

import (
	"letsplay-microservice/internal/client"
	"letsplay-microservice/internal/config"
	"letsplay-microservice/internal/handler"
	"letsplay-microservice/internal/service"

	"go.uber.org/zap"
)

type Container struct {
	PlayerHandler *handler.PlayerHandler
}

func BuildContainer(config *config.Config, logger *zap.Logger) *Container {
	playerClient := client.NewPlayerClient(config.AuthServerUrl, logger)
	playerService := service.NewPlayerService(playerClient)
	playerHandler := handler.NewPlayerHandler(playerService)

	return &Container{
		PlayerHandler: playerHandler,
	}
}
