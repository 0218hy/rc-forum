package authhttp

import (
	"context"
	repo "rc-forum-backend/db/sqlc"
)

type Service interface {
	CreateSession(ctx context.Context, params CreateSessionParams) (*repo.Session, error)
	RevokeSession(ctx context.Context, sessionID string) error
	DeleteSessionsByUserID(ctx context.Context, userID int32) error
	GetSession(ctx context.Context, sessionID string) (*repo.Session, error)
}

type svc struct {
	repo *repo.Queries
}

func NewService(repo *repo.Queries) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) CreateSession(ctx context.Context, params CreateSessionParams) (*repo.Session, error) {
	createdSession, err := s.repo.CreateSession(ctx, repo.CreateSessionParams{
		UserID:       params.UserID,
		RefreshToken: params.RefreshToken,
		IsRevoked:    params.IsRevoked,
		ExpiresAt:    params.ExpiresAt,
	})
	if err != nil {
		return nil, err
	}

	return &createdSession, nil
}

func (s *svc) RevokeSession(ctx context.Context, sessionID string) error {
	return s.repo.RevokeSession(ctx, sessionID)
}

func (s *svc) DeleteSessionsByUserID(ctx context.Context, userID int32) error {
	return s.repo.DeleteSessionsByUserID(ctx, userID)
}

func (s *svc) GetSession(ctx context.Context, sessionID string) (*repo.Session, error) {
	session, err := s.repo.GetSession(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	return &session, nil
}