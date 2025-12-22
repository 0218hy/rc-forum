package users

import (
	"context"
	repo "rc-forum-backend/db/sqlc"
)

type Service interface {
	GetUserProfile(ctx context.Context, userID int64) (repo.User, error)
}

type svc struct {
	repo *repo.Queries
}

func NewService(repo *repo.Queries) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) GetUserProfile(ctx context.Context, userID int64) (repo.User, error) {
	return s.repo.GetUserByID(ctx, userID)
}