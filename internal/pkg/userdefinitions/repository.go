package userdefinitions

import (
	"cyclolab-microservice/internal/model"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(userID uuid.UUID, definitions model.UserDefinitions) error {
	query := `
		INSERT INTO user_definitions (user_id, given_name, surname, birthdate, weight, height, gender)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(query,
		userID,
		definitions.GivenName,
		definitions.Surname,
		definitions.Birthdate,
		definitions.Weight,
		definitions.Height,
		definitions.Gender,
	)

	return err
}

func (r *Repository) Get(userID uuid.UUID) (*model.UserDefinitions, error) {
	var (
		givenName, surname string
		weight             float32
		height             float32
		gender             string
		birthdate          time.Time
	)

	query := `
		SELECT given_name, surname, birthdate, weight, height, gender
		FROM user_definitions
		WHERE user_id = $1
	`

	err := r.db.QueryRow(query, userID).Scan(
		&givenName,
		&surname,
		&birthdate,
		&weight,
		&height,
		&gender,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &model.UserDefinitions{
		UserID:    userID,
		Birthdate: birthdate,
		GivenName: givenName,
		Surname:   surname,
		Weight:    weight,
		Height:    height,
		Gender:    gender,
	}, nil
}

func (r *Repository) Upsert(userID uuid.UUID, definitions model.UserDefinitions) error {
	query := `
		INSERT INTO user_definitions (user_id, given_name, surname, birthdate, gender)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id) DO UPDATE SET
			given_name = EXCLUDED.given_name,
			surname = EXCLUDED.surname,
			birthdate = EXCLUDED.birthdate,
			gender = EXCLUDED.gender
	`

	_, err := r.db.Exec(query,
		userID,
		definitions.GivenName,
		definitions.Surname,
		definitions.Birthdate,
		definitions.Gender,
	)

	return err
}
