package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
)

// Connect opens a new connection to a Postgres server.
func Connect() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
}
