package authhttp

import (
	repo "rc-forum-backend/db/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func toUser(u repo.User) User {
	return User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		IsAdmin:  u.IsAdmin,
	}
}	

type RegisterUserPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateSessionParams struct {
	ID           int32            `json:"id"`
	UserID       int32            `json:"user_id"`
	RefreshToken string           `json:"refresh_token"`
	IsRevoked    bool             `json:"is_revoked"`
	ExpiresAt    pgtype.Timestamp `json:"expires_at"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
}

type SessionResponse struct {
	SessionID               string  `json:"session_id"`
	AccessToken             string  `json:"access_token"`
	RefreshToken            string  `json:"refresh_token"`
	AccessTokenExpiresAt    time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt   time.Time `json:"refresh_token_expires_at"`
	User                    User      `json:"user"`
}

type RenewAccessTokenPayload struct {
	RefreshToken string `json:"refresh_token"`
}

type RenewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}