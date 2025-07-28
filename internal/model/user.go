package model

import "time"

type UserAuth struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserDefinitions struct {
	Nickname       string     `json:"nickname"`
	Birthdate      time.Time  `json:"birthdate"`
	PreferredSport GameInfo   `json:"preferredSport"`
	OtherSports    []GameInfo `json:"otherSports"`
}
