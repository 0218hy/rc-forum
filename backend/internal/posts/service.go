package posts

import (
	"context"
	repo "rc-forum-backend/db/sqlc"
)

type Service interface {
	GetAllPosts(ctx context.Context) ([]repo.GetAllPostsRow, error)

}

type svc struct {
	repo *repo.Queries
}

func NewService(repo *repo.Queries) Service {
	return &svc {
		repo: repo,
	}
}

func (s *svc) GetAllPosts(ctx context.Context) ([]repo.GetAllPostsRow, error) {
	return s.repo.GetAllPosts(ctx)
}