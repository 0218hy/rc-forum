package users

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"rc-forum-backend/internal/json"
	"strconv"

	"github.com/go-chi/chi/v5"
)


type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

// AUTH: FUNC GetMyProfile

func (h *handler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	// 1. Call service -> GetUserProfile
	id := chi.URLParam(r, "id")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	profile, err := h.service.GetUserProfile(r.Context(), userID)
	if err != nil {
		 if errors.Is(err, sql.ErrNoRows) {
			log.Println("No user found with that ID")
			http.Error(w, "User not found", http.StatusNotFound);
			return
		} else {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// 2. Return JSON response into an HTTP response
	json.Write(w, http.StatusOK, profile)
}
