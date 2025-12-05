package bootstrap

import (
	"cyclolab-microservice/internal/client"
	"cyclolab-microservice/internal/config"
	"cyclolab-microservice/internal/database"
	"cyclolab-microservice/internal/handler"
	"cyclolab-microservice/internal/pkg/settings"
	"cyclolab-microservice/internal/pkg/userdefinitions"
	"cyclolab-microservice/internal/service"

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
