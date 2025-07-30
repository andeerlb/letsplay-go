package model

type Settings struct {
	Layout   string `json:"layout" db:"layout"`
	Language string `json:"language" db:"language"`
}
