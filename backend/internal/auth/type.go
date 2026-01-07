package auth

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	ID       int32 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}



