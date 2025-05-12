package read_model

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetConnection() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/marketdb")
}
