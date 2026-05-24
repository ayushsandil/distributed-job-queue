package db

import (
    "context"
    "os"

    "github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() error {
    databaseUrl := os.Getenv("POSTGRES_URL")

    pool, err := pgxpool.New(context.Background(), databaseUrl)
    if err != nil {
        return err
    }

    DB = pool
    return nil
}
