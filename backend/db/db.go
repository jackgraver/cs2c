package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect() {
	conn, err := pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/mydb")
	if err != nil {
		return
	}
	defer conn.Close(context.Background())
}
