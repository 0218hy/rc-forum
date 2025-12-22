package users

import (
	"context"
	repo "rc-forum-backend/db/sqlc"

	"github.com/jackc/pgx/v5"
)

type Service interface {
	GetUserProfile(ctx context.Context, userID int64) (repo.User, error)
}

type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) GetUserProfile(ctx context.Context, userID int64) (repo.User, error) {
	return s.repo.GetUserByID(ctx, userID)
}