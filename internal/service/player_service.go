package service

import (
	"context"
	"letsplay-microservice/internal/client"
	"letsplay-microservice/internal/model"
	"letsplay-microservice/internal/pkg/userdefinitions"
)

type PlayerService struct {
	client     *client.PlayerClient
	repository *userdefinitions.Repository
}

func NewPlayerService(client *client.PlayerClient, repo *userdefinitions.Repository) *PlayerService {
	return &PlayerService{
		client:     client,
		repository: repo,
	}
}

func (ps *PlayerService) CreateNewPlayer(ctx context.Context, payload model.SignUp) (*model.AuthTokenResponse, error) {
	player, err := ps.client.CreatePlayer(ctx, payload.UserAuth)

	if err != nil {
		return nil, err
	}

	err = ps.repository.Save(player.User.Id, payload.UserDefinitions)

	if err != nil {
		return nil, err
	}
	
	return player, nil
}
