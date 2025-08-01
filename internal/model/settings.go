package model

type Settings struct {
	Layout   string `json:"layout" db:"layout" binding:"required"`
	Language string `json:"language" db:"language" binding:"required"`
}
