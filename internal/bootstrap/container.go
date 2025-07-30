package bootstrap

import (
	"letsplay-microservice/internal/client"
	"letsplay-microservice/internal/config"
	"letsplay-microservice/internal/database"
	"letsplay-microservice/internal/handler"
	"letsplay-microservice/internal/pkg/settings"
	"letsplay-microservice/internal/pkg/userdefinitions"
	"letsplay-microservice/internal/service"

	"go.uber.org/zap"
)

type Container struct {
	UserHandler    *handler.UserHandler
	SettingHandler *handler.SettingsHandlers
}

func BuildContainer(config *config.Config, logger *zap.Logger) *Container {
	playerClient := client.NewPlayerClient(config.AuthServerUrl, logger)

	userDefinitionsRepo := userdefinitions.NewRepository(database.DB)
	settingsRepo := settings.NewRepository(database.DB)

	playerService := service.NewUserService(playerClient, userDefinitionsRepo)
	settingsService := service.NewSettingsService(settingsRepo)

	// Handler
	userHandler := handler.NewUserHandler(playerService)
	settingsHandler := handler.NewSettingsHandler(settingsService)

	return &Container{
		UserHandler:    userHandler,
		SettingHandler: settingsHandler,
	}
}
