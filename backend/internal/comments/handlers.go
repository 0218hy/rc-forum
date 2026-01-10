package comments

import (
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

func (h *handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	postID, err := utility.GetPostID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims, ok := auth.FromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req CreateCommentRequest
	if err := json.Read(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment, err := h.service.CreateComment(r.Context(), CreateCommentRequest{
		PostID:   postID,
		AuthorID: claims.ID,
		Body:     req.Body,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, comment)
}


func (h *handler) ListCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	postID, err := utility.GetPostID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comments, err := h.service.ListCommentsByPostID(r.Context(), postID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, comments)
}

func (h *handler) DeleteCommentByID(w http.ResponseWriter, r *http.Request) {
	commentID, err := utility.GetCommentID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims, ok := auth.FromContext(r.Context())
	if !ok {
		log.Println("Failed to get author ID from context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	authorID := claims.ID

	err = h.service.DeleteCommentByID(r.Context(), commentID, authorID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, nil)
}
