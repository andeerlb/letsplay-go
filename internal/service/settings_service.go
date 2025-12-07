package service

import (
	"context"
	"cyclolab-microservice/internal/middleware"
	"cyclolab-microservice/internal/model"
	"cyclolab-microservice/internal/pkg/settings"
	"fmt"

	"github.com/google/uuid"
)

type SettingsService struct {
	repository *settings.Repository
}

func NewSettingsService(repo *settings.Repository) *SettingsService {
	return &SettingsService{
		repository: repo,
	}
}

func (us *SettingsService) Update(ctx context.Context, settings model.Settings) (*model.Settings, error) {
	userUUID, _ := ctx.Value(middleware.UserIDKey).(uuid.UUID)
	userSettings, err := us.repository.Upsert(userUUID, settings)
	if err != nil {
		return nil, fmt.Errorf("FAILED_TO_UPDATE_SETTINGS")
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
