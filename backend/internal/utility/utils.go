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

func GetPostID (r *http.Request) (int32, error) {
	postIDStr := chi.URLParam(r, "post_id")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil || postID <= 0 {
		return 0, errors.New("Invalid post ID")		
	}
	return int32(postID), nil
}

func GetCommentID (r *http.Request) (int32, error) {
	commentIDStr := chi.URLParam(r, "comment_id")
	commentID, err := strconv.ParseInt(commentIDStr, 10, 64)
	if err != nil || commentID <= 0 {
		return 0, errors.New("Invalid comment ID")		
	}
	return int32(commentID), nil
}

func GetSessionID(r *http.Request) (int32, error) {
	sessionIDStr := chi.URLParam(r, "session_id")
	sessionID, err := strconv.ParseInt(sessionIDStr, 10, 64)
	if err != nil || sessionID <= 0 {
		return 0, errors.New("Invalid session ID")		
	}
	return int32(sessionID), nil
}

