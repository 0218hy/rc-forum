package products

import (
	"context"
	repo "rc-forum-backend/db/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) (error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) (error) {
	return nil
}
