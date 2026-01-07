package utility

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)


func GetID(r *http.Request) (int32, error) {
	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil || id <= 0 {
		return 0, errors.New("Invalid ID")
	}
	return int32(id), nil
}

func GetEmail(r *http.Request) (string, error) {
	email := chi.URLParam(r, "email")
	if email == "" {
		return "", errors.New("Email is required")
	}

	return email, nil
}

func GetSessionID(r *http.Request) (int32, error) {
	sessionIDStr := chi.URLParam(r, "session_id")
	sessionID, err := strconv.ParseInt(sessionIDStr, 10, 64)
	if err != nil || sessionID <= 0 {
		return 0, errors.New("Invalid session ID")		
	}
	return int32(sessionID), nil
}

func IsPostOwner(postAuthorID int32, requesterID int32) bool {
	return postAuthorID == requesterID
}