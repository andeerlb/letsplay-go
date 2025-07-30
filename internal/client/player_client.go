package client

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"os"
	"time"

	"letsplay-microservice/internal/model"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type PlayerClient struct {
	baseURL string
	client  *resty.Client
	logger  *zap.Logger
}

func NewPlayerClient(baseURL string, logger *zap.Logger) *PlayerClient {
	return &PlayerClient{
		baseURL: baseURL,
		client:  NewRestyClientWithRetry(5*time.Second, 3),
		logger:  logger,
	}
}

func (pc *PlayerClient) CreateUserAccount(ctx context.Context, payload model.UserAuth) (*model.AuthTokenResponse, error) {
	var created model.AuthTokenResponse

	resp, err := pc.client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(&created).
		Post(pc.baseURL + "/signup")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		pc.logger.Debug("create user error",
			zap.Int("status_code", resp.StatusCode()),
			zap.String("status", resp.Status()),
			zap.String("body", resp.String()),
		)

		if resp.StatusCode() == 400 {
			return nil, errors.New("USER_ALREADY_EXISTS")
		}
		return nil, errors.New("FAILED_TO_CREATE_USER")
	}

	return &created, nil
}

func (pc *PlayerClient) DeleteUser(userUUID uuid.UUID, ctx context.Context) (*bool, error) {
	resp, err := pc.client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+os.Getenv("LETSPLAY_JWT_ADMIN_TOKEN")).
		Delete(pc.baseURL + "/admin/users/" + userUUID.String())

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() > 299 {
		return nil, errors.New("FAILED_TO_DELETE_USER")
	}

	success := true
	return &success, nil
}
