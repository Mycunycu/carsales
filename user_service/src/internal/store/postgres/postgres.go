package postgres

import (
	"carsalesuser/internal/config"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PgStore struct {
	*pgxpool.Pool
}

func Connect(cfg *config.Config) (*PgStore, error) {
	pool, err := pgxpool.Connect(context.Background(), cfg.PgConnection)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	return &PgStore{pool}, nil
}
