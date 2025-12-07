package settings

import (
	"cyclolab-microservice/internal/model"
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Upsert(userID uuid.UUID, settings model.Settings) (*model.Settings, error) {
	var columns []string
	var placeholders []string
	var values []any
	var updates []string

	paramIndex := 1

	if settings.Layout != "" {
		columns = append(columns, "layout")
		placeholders = append(placeholders, fmt.Sprintf("$%d", paramIndex))
		values = append(values, settings.Layout)
		updates = append(updates, "layout = EXCLUDED.layout")
		paramIndex++
	}

	if settings.Language != "" {
		columns = append(columns, "language")
		placeholders = append(placeholders, fmt.Sprintf("$%d", paramIndex))
		values = append(values, settings.Language)
		updates = append(updates, "language = EXCLUDED.language")
		paramIndex++
	}

	if len(columns) == 0 {
		return nil, nil
	}

	columns = append([]string{"user_id"}, columns...)
	placeholders = append([]string{fmt.Sprintf("$%d", paramIndex)}, placeholders...)
	values = append(values, userID)

	query := fmt.Sprintf(`
		INSERT INTO settings (%s)
		VALUES (%s)
		ON CONFLICT (user_id) DO UPDATE
		SET %s
		RETURNING layout, language
	`,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
		strings.Join(updates, ", "),
	)

	var layout, language sql.NullString
	err := r.db.QueryRow(query, values...).Scan(&layout, &language)
	if err != nil {
		return nil, err
	}

	return &model.Settings{
		Layout:   layout.String,
		Language: language.String,
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
