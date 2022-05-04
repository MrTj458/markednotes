package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Connect opens a new connection to a Postgres server.
func Connect() (*pgxpool.Pool, error) {
	return pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
}
