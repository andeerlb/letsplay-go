package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"letsplay-microservice/internal/config"
	"log"
)

var DB *sqlx.DB

func InitPostgres(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.SSLMode,
		cfg.DB.SearchPath,
	)

	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Could not connect to PostgreSQL: %v", err)
	}

	log.Println("Connected to PostgreSQL")
}
