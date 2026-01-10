package comments

import (
	"context"
	repo "rc-forum-backend/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service interface {
	CreateComment(ctx context.Context, req CreateCommentRequest) (repo.Comment, error)
	DeleteCommentByID(ctx context.Context, commentID int32, authorID int32) error
	ListCommentsByPostID(ctx context.Context, postID int32) ([]repo.Comment, error)
}

type svc struct {
	repo *repo.Queries
	db   *pgxpool.Pool
}

func NewService(repo *repo.Queries, db *pgxpool.Pool) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) CreateComment(ctx context.Context, req CreateCommentRequest) (repo.Comment, error) {
	comment, err := s.repo.CreateComment(ctx, repo.CreateCommentParams{
		PostID:   req.PostID,
		AuthorID: req.AuthorID,
		Body:     req.Body,
	})
	if err != nil {
		return repo.Comment{}, err
	}

	return comment, nil
}

func (s *svc) DeleteCommentByID(ctx context.Context, commentID int32, authorID int32) error {
	return s.repo.DeleteCommentByID(ctx, repo.DeleteCommentByIDParams{
		ID:       commentID,
		AuthorID: authorID,
	})
}

func (s *svc) ListCommentsByPostID(ctx context.Context, postID int32) ([]repo.Comment, error) {
	return s.repo.ListCommentsByPostID(ctx, postID)
}
