package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func customConfig() (*pgxpool.Config, error) {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	config.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {
		c.TypeMap().RegisterDefaultPgType(pgx.TextFormatCode, "text")
		return nil
	}

	return config, nil
}

func SetupDbConnection() (*pgxpool.Pool, error) {
	pgxConfig, err := customConfig()
	dbpool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		return nil, err
	}

	return dbpool, err
}
