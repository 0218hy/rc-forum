package posts

import (
	"log"
	"net/http"
	"rc-forum-backend/internal/json"

)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler {
		service: service,
	}
}

func (h *handler) GetAllPosts(w http.ResponseWriter, r *http.Request){
	// 1. Call servicce -> GetAllPosts
	allposts, err := h.service.GetAllPosts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 2 Return JSON response into an HTTP response
	json.Write(w, http.StatusOK, allposts)
}