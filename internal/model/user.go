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
	UserID         uuid.UUID  `db:"user_id"`
	Nickname       string     `json:"nickname" db:"nickname"`
	Birthdate      time.Time  `json:"birthdate" db:"birthdate"`
	PreferredSport GameInfo   `json:"preferredSport" db:"preferred_sport"`
	OtherSports    []GameInfo `json:"otherSports" db:"other_sports"`
}
