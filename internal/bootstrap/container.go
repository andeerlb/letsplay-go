package bootstrap

import (
	"letsplay-microservice/internal/client"
	"letsplay-microservice/internal/config"
	"letsplay-microservice/internal/database"
	"letsplay-microservice/internal/handler"
	"letsplay-microservice/internal/pkg/userdefinitions"
	"letsplay-microservice/internal/service"

	"go.uber.org/zap"
)

type Container struct {
	UserHandler *handler.UserHandler
}

func BuildContainer(config *config.Config, logger *zap.Logger) *Container {
	playerClient := client.NewPlayerClient(config.AuthServerUrl, logger)

	userDefinitionsRepo := userdefinitions.NewRepository(database.DB)

	playerService := service.NewUserService(playerClient, userDefinitionsRepo)

	// Handler
	userHandler := handler.NewUserHandler(playerService)

	return &Container{
		UserHandler: userHandler,
	}
}
