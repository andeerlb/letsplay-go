package settings

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"letsplay-microservice/internal/model"
	"strings"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(userID uuid.UUID, settings model.Settings) (*model.Settings, error) {
	fields := []string{"user_id"}
	values := []any{userID}
	placeholders := []string{"$1"}

	var updateFields []string
	paramIndex := 2

	if settings.Layout != "" {
		fields = append(fields, "layout")
		values = append(values, settings.Layout)
		placeholders = append(placeholders, fmt.Sprintf("$%d", paramIndex))
		updateFields = append(updateFields, fmt.Sprintf("layout = EXCLUDED.layout"))
		paramIndex++
	}

	if settings.Language != "" {
		fields = append(fields, "language")
		values = append(values, settings.Language)
		placeholders = append(placeholders, fmt.Sprintf("$%d", paramIndex))
		updateFields = append(updateFields, fmt.Sprintf("language = EXCLUDED.language"))
		paramIndex++
	}

	if len(fields) == 1 {
		return nil, nil
	}

	query := fmt.Sprintf(`
		INSERT INTO settings (%s)
		VALUES (%s)
		ON CONFLICT (user_id) DO UPDATE SET %s
		RETURNING layout, language
	`, strings.Join(fields, ", "), strings.Join(placeholders, ", "), strings.Join(updateFields, ", "))

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
