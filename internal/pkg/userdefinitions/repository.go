package userdefinitions

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"letsplay-microservice/internal/model"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(userID uuid.UUID, definitions model.UserDefinitions) error {
	preferredSportJSON, err := json.Marshal(definitions.PreferredSport)
	if err != nil {
		return err
	}

	otherSportsJSON, err := json.Marshal(definitions.OtherSports)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO user_definitions (user_id, nickname, birthdate, preferred_sport, other_sports)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id) DO UPDATE SET
			nickname = EXCLUDED.nickname,
			birthdate = EXCLUDED.birthdate,
			preferred_sport = EXCLUDED.preferred_sport,
			other_sports = EXCLUDED.other_sports
	`

	_, err = r.db.Exec(query,
		userID,
		definitions.Nickname,
		definitions.Birthdate,
		preferredSportJSON,
		otherSportsJSON,
	)

	return err
}
