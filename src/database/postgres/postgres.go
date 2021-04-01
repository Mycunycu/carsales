package postgres

import (
	"carsales/internal/config"
	"database/sql"

	_ "github.com/lib/pq"
)

type PgStore struct {
	*sql.DB
}

func Connect(cfg *config.Config) (*PgStore, error) {
	db, err := sql.Open("postgres", cfg.PgConnStr)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	return &PgStore{db}, nil
}
