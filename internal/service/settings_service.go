package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"letsplay-microservice/internal/middleware"
	"letsplay-microservice/internal/model"
	"letsplay-microservice/internal/pkg/settings"
)

type SettingsService struct {
	repository *settings.Repository
}

func NewSettingsService(repo *settings.Repository) *SettingsService {
	return &SettingsService{
		repository: repo,
	}
}

func (us *SettingsService) Save(ctx context.Context, settings model.Settings) (*model.Settings, error) {
	userUUID, _ := ctx.Value(middleware.UserIDKey).(uuid.UUID)
	userSettings, err := us.repository.Save(userUUID, settings)
	if err != nil {
		return nil, fmt.Errorf("FAILED_TO_SAVE_SETTINGS")
	}
	return userSettings, nil
}

func (us *SettingsService) Get(ctx context.Context) (*model.Settings, error) {
	userUUID, _ := ctx.Value(middleware.UserIDKey).(uuid.UUID)
	fetchSettings, err := us.repository.Get(userUUID)
	if err != nil {
		return nil, fmt.Errorf("FAILED_TO_GET_SETTINGS")
	}
	return fetchSettings, nil
}
