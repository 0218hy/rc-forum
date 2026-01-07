package users

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"rc-forum-backend/internal/auth"
	"rc-forum-backend/internal/json"
	"rc-forum-backend/internal/utility"
)


type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetMyProfile(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(auth.AuthKey{}).(*auth.UserClaims)

	user, err := h.service.GetUserByID(r.Context(), claims.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get user profile", http.StatusInternalServerError)
		return
	}
	
	json.Write(w, http.StatusOK, user)
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// 1. Call service -> GetUserProfile
	userID, err := utility.GetID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(r.Context(), userID)
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
	json.Write(w, http.StatusOK, user)
}

func (h *handler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email, err := utility.GetEmail(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	user, err := h.service.GetUserByEmail(r.Context(), email)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid user email", http.StatusBadRequest)
		return
	}

	json.Write(w, http.StatusOK, user)
}
