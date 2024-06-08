package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupDbConnection() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return dbpool, err
}