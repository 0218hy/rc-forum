package users

import (
	"context"
	repo "rc-forum-backend/db/sqlc"
)

type Service interface {
	GetUserByID(ctx context.Context, userID int32) (repo.User, error)
	GetUserByEmail(ctx context.Context, email string) (repo.User, error)
	CreateUser(ctx context.Context, tempUser CreateUserParams) (int32, error)
}

type svc struct {
	repo *repo.Queries
}

func NewService(repo *repo.Queries) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) GetUserByID(ctx context.Context, userID int32) (repo.User, error) {
	return s.repo.FindUserByID(ctx, userID)
}

func (s *svc) GetUserByEmail(ctx context.Context, email string) (repo.User, error) {
	return s.repo.FindUserByEmail(ctx, email)
}

func (s *svc) CreateUser(ctx context.Context, tempUser CreateUserParams) (int32, error) {
	userID, err := s.repo.CreateUser(ctx, repo.CreateUserParams{
		Name: tempUser.Name,
		Email: tempUser.Email,
		Password: tempUser.Password,
		IsAdmin: tempUser.IsAdmin,
	})
	if err != nil {
		return 0, err
	}

	return userID, nil
}