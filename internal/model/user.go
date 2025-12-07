package model

import (
	"time"

	"github.com/google/uuid"
)

type UserAuth struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserDefinitions struct {
	UserID    uuid.UUID `json:"-" db:"user_id"`
	GivenName string    `json:"givenName" db:"given_name" binding:"required"`
	Surname   string    `json:"surname" db:"surname" binding:"required"`
	Birthdate time.Time `json:"birthdate" db:"birthdate" binding:"required"`
	Weight    float32   `json:"weight" db:"weight" binding:"required"`
	Height    float32   `json:"height" db:"height" binding:"required"`
	Gender    string    `json:"gender" db:"gender" binding:"required,oneof=M F"`
}

type UserDefinitionsUpdateRequest struct {
	UserID    uuid.UUID `json:"-" db:"user_id"`
	GivenName string    `json:"givenName" db:"given_name" binding:"required"`
	Surname   string    `json:"surname" db:"surname" binding:"required"`
	Birthdate time.Time `json:"birthdate" db:"birthdate" binding:"required"`
	Gender    string    `json:"gender" db:"gender" binding:"required,oneof=M F"`
}
