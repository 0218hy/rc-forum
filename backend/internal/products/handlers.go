package products

import (
	"log"
	"net/http"
	"rc-forum-backend/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1. Call service -> ListProducts
	err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 2. Return JSON response into an HTTP response
	products := []string{"Product A", "Product B", "Product C"}

	json.Write(w, http.StatusOK, products)
}