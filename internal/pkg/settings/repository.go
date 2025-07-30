package settings

import (
	"database/sql"
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

func (r *Repository) Save(userID uuid.UUID, settings model.Settings) (*model.Settings, error) {
	var (
		layout   string
		language string
	)

	query := `
		INSERT INTO settings (user_id, layout, language)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id) DO UPDATE SET
			layout = EXCLUDED.nickname,
			language = EXCLUDED.birthdate
	`

	r.db.QueryRow(query, userID, settings.Layout, settings.Language).Scan(
		&layout,
		&language,
	)

	return &model.Settings{
		Layout:   layout,
		Language: language,
	}, nil
}

func (r *Repository) Get(userID uuid.UUID) (*model.Settings, error) {
	var (
		layout   string
		language string
	)

	query := `
		SELECT layout, language
		FROM settings
		WHERE user_id = $1
	`

	err := r.db.QueryRow(query, userID).Scan(
		&layout,
		&language,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &model.Settings{
		Layout:   layout,
		Language: language,
	}, nil
}
