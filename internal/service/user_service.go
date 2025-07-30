package service

import (
	"context"
	"fmt"
	"letsplay-microservice/internal/client"
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

	if err := us.repository.Save(player.User.Id, payload.UserDefinitions); err != nil {
		_, _ = us.client.DeleteUser(player.User.Id, ctx)
		return nil, fmt.Errorf("FAILED_TO_DELETE_USER")
	}

	return player, nil
}
