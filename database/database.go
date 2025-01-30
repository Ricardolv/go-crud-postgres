package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func NewPostgresConnection(connString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao PostgreSQL: %v", err)
	}
	return conn, nil
}
