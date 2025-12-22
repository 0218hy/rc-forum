package repository

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB() (*pgxpool.Pool, error) {
	conn := "postgresql://rc_user:rc_password@localhost:5432/rc_forum"
	return pgxpool.New(context.Background(), conn)
}
