package comments

import (
	"net/http"

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

func (h *handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	
}