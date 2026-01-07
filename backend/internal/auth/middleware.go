package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"rc-forum-backend/internal/json"
	"strings"
)

type AuthKey struct{} 

func GetAuthMiddlewareFunc(tokenMaker *JWTMaker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. Get the authorization header & verify the token
			claims, err := verifyClaimsFromAuthHeader(r, tokenMaker)
			if err != nil {
				log.Println("Auth middleware - verify token error:", err)
				json.Write(w, http.StatusUnauthorized, "error: unauthorized")
				return
			}
			// 2. Pass the claims down the context
			ctx := context.WithValue(r.Context(), AuthKey{}, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetAdminMiddlewareFunc(tokenMaker *JWTMaker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. Get the authorization header & verify the token
			claims, err := verifyClaimsFromAuthHeader(r, tokenMaker)
			if err != nil {
				log.Println("Admin middleware - verify token error:", err)
				json.Write(w, http.StatusUnauthorized, "error: unauthorized")
				return
			}
			// 2. Check if user is admin
			if !claims.IsAdmin {
				log.Println("Admin middleware - user is not admin")
				json.Write(w, http.StatusForbidden, "error: forbidden")
				return
			}
			// 3. Pass the claims down the context
			ctx := context.WithValue(r.Context(), AuthKey{}, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func verifyClaimsFromAuthHeader(r *http.Request, tokenMaker *JWTMaker) (*UserClaims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is missing")
	}

	// Slice authHeader (eg. "Bearer <token>")
	fields := strings.Fields(authHeader)
	if len(fields) != 2 || strings.ToLower(fields[0]) != "bearer" {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	accessToken := fields[1]
	claims, err := tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return claims, nil
}