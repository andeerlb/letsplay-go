package model

import (
	"github.com/google/uuid"
	"time"
)

type UserAuth struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserDefinitions struct {
	UserID         uuid.UUID  `json:"-" db:"user_id"`
	Nickname       string     `json:"nickname" db:"nickname" binding:"required"`
	Birthdate      time.Time  `json:"birthdate" db:"birthdate" binding:"required"`
	PreferredSport GameInfo   `json:"preferredSport" db:"preferred_sport" binding:"required"`
	OtherSports    []GameInfo `json:"otherSports,omitempty" db:"other_sports"`
}
