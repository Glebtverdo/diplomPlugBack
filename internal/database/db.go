package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func InitDatabase() error {
	var err error
	pool, err = pgxpool.New(context.Background(), "host=localhost port=5432 dbname=diplomName user=postgres sslmode=prefer connect_timeout=10")
	if err != nil {
		return err
	}
	return nil
}
