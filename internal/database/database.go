package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDBConn() *pgxpool.Pool {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	return pool
}
