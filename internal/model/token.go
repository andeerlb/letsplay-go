package model

import "github.com/google/uuid"

type AuthTokenUser struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

type AuthTokenResponse struct {
	AccessToken  string        `json:"access_token"`
	TokenType    string        `json:"token_type"`
	RefreshToken string        `json:"refresh_token"`
	ExpiresIn    int           `json:"expires_in"`
	ExpiresAt    int64         `json:"expires_at"`
	User         AuthTokenUser `json:"user"`
}
