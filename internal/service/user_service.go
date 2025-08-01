package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"letsplay-microservice/internal/client"
	"letsplay-microservice/internal/locale"
	"letsplay-microservice/internal/middleware"
	"letsplay-microservice/internal/model"
	"letsplay-microservice/internal/pkg/userdefinitions"
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
