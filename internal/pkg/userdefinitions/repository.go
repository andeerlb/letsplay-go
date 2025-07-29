package userdefinitions

import (
	"github.com/jmoiron/sqlx"
	"letsplay-microservice/internal/model"
)

func SaveUserDefinitions(db *sqlx.DB, definitions model.UserDefinitions) error {
	query := `
		INSERT INTO user_definitions (user_id, nickname, birthdate, preferred_sport, other_sports)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id) DO UPDATE SET
			nickname = EXCLUDED.nickname,
			birthdate = EXCLUDED.birthdate,
			preferred_sport = EXCLUDED.preferred_sport,
			other_sports = EXCLUDED.other_sports
	`

	_, err := db.Exec(query,
		definitions.UserID,
		definitions.Nickname,
		definitions.Birthdate,
		definitions.PreferredSport,
		definitions.OtherSports,
	)

	return err
}
