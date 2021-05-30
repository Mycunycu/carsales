package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type PgStore struct {
	*pgxpool.Pool
}

func Connect(dbUrl string) (*PgStore, error) {
	pool, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}

	return &PgStore{pool}, nil
}

func MigrateDatabase(source string, dbUrl string) error {
	m, err := migrate.New(source, dbUrl)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	return nil
}
