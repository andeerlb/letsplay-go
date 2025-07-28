package service

import (
	"context"
	"letsplay-microservice/internal/client"
	"letsplay-microservice/internal/model"
)

type PlayerService struct {
	client *client.PlayerClient
}

func NewPlayerService(client *client.PlayerClient) *PlayerService {
	return &PlayerService{client: client}
}

func (ps *PlayerService) CreateNewPlayer(ctx context.Context, payload model.SignUpAuthServer) (*model.AuthTokenResponse, error) {
	return ps.client.CreatePlayer(ctx, payload)
}
