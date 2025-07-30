package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"letsplay-microservice/internal/client"
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
	if payload.UserDefinitions.Nickname == "" {
		return nil, fmt.Errorf("FAILED_TO_DELETE_USER_REQUIRED_FIELDS_IS_EMPTY: Nickname is empty")
	}

	if payload.UserDefinitions.Birthdate.IsZero() {
		return nil, fmt.Errorf("FAILED_TO_DELETE_USER_REQUIRED_FIELDS_IS_EMPTY: Birthdate is empty")
	}

	if payload.UserDefinitions.PreferredSport.Position == "" {
		return nil, fmt.Errorf("FAILED_TO_DELETE_USER_REQUIRED_FIELDS_IS_EMPTY: PreferredSport.Position is empty")
	}

	if payload.UserDefinitions.PreferredSport.Type == "" {
		return nil, fmt.Errorf("FAILED_TO_DELETE_USER_REQUIRED_FIELDS_IS_EMPTY: PreferredSport.Type is empty")
	}

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

func (us *UserService) GetUserDefinitions(ctx context.Context) (*model.UserDefinitions, error) {
	userUUID, _ := ctx.Value(middleware.UserIDKey).(uuid.UUID)
	userDefinitions, err := us.repository.Get(userUUID)
	if err != nil {
		return nil, fmt.Errorf("FAILED_TO_GET_USER_DEFINITIONS")
	}
	return userDefinitions, nil
}
