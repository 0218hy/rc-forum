package posts

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"rc-forum-backend/internal/json"
	"rc-forum-backend/internal/utility"
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

func (h *handler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	postID, err := utility.GetID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := h.service.GetPostByID(r.Context(), postID)
	if err != nil {
		 if errors.Is(err, sql.ErrNoRows) {
			log.Println("No post found with that ID")
			http.Error(w, "Post not found", http.StatusNotFound);
			return
		} else {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	json.Write(w, http.StatusOK, post)
}

func (h *handler) DeletePostByID (w http.ResponseWriter, r *http.Request) {
	postID, err := utility.GetID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeletePostByID(r.Context(), postID)
	if err != nil {
		 if errors.Is(err, sql.ErrNoRows) {
			log.Println("No post found with that ID")
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		} else {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *handler) CreatePost (w http.ResponseWriter, r *http.Request){
	var tempPost CreatePostRequest
	if err := json.Read(r, &tempPost); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdPost, err := h.service.CreatePost(r.Context(), tempPost)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	json.Write(w, http.StatusCreated, createdPost)
}


func (h *handler) UpdatePostCore (w http.ResponseWriter, r *http.Request){
	id , err := utility.GetID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var tempPost UpdatePostCoreRequest
	if err := json.Read(r, &tempPost); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdatePostCore(r.Context(), id, tempPost); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, nil)
}