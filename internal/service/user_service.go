package service

import (
	"context"
	"cyclolab-microservice/internal/client"
	"cyclolab-microservice/internal/locale"
	"cyclolab-microservice/internal/middleware"
	"cyclolab-microservice/internal/model"
	"cyclolab-microservice/internal/pkg/userdefinitions"
	"fmt"

	"github.com/google/uuid"
)

type UserService struct {
	client     *client.PlayerClient
	repository *userdefinitions.Repository
}

func NewUserService(client *client.PlayerClient, repo *userdefinitions.Repository) *UserService {
	return &UserService{
		client:     client,
		repository: repo,
	}
}

func (us *UserService) SignUp(ctx context.Context, payload model.SignUp) (*model.AuthTokenResponse, error) {
	player, err := us.client.CreateUserAccount(ctx, payload.UserAuth)
	if err != nil {
		return nil, err
	}

	msg := locale.Msg(ctx, "user_service.failed-to-delete-user")

	if err := us.repository.Save(player.User.Id, payload.UserDefinitions); err != nil {
		_, _ = us.client.DeleteUser(player.User.Id, ctx)
		return nil, fmt.Errorf(msg)
	}

	return player, nil
}

func (us *UserService) GetUserDefinitions(ctx context.Context) (*model.UserDefinitions, error) {
	userUUID, _ := ctx.Value(middleware.UserIDKey).(uuid.UUID)
	userDefinitions, err := us.repository.Get(userUUID)
	if err != nil {
		return nil, fmt.Errorf("FAILED_TO_GET_USER_DEFINITIONS")
	}
	return userDefinitions, nil
}

func (us *UserService) UpdateUserDefinitions(ctx context.Context, definitions model.UserDefinitions) (*model.UserDefinitions, error) {
	userUUID, _ := ctx.Value(middleware.UserIDKey).(uuid.UUID)
	err := us.repository.Upsert(userUUID, definitions)
	if err != nil {
		return nil, fmt.Errorf("FAILED_TO_UPDATE_USER_DEFINITIONS")
	}
	return &definitions, nil
}
