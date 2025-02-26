package handlers

import (
	"encoding/json"
	"net/http"

	"microservices-travel-backend/internal/blog-service/domain/models"
	"microservices-travel-backend/internal/blog-service/domain/ports"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type BlogHandler struct {
	blogService ports.BlogServicePort
}

func NewBlogHandler(service ports.BlogServicePort) *BlogHandler {
	return &BlogHandler{blogService: service}
}

func (h *BlogHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/blogs", h.CreateBlog).Methods(http.MethodPost)
	router.HandleFunc("/blogs/{id}", h.GetBlogByID).Methods(http.MethodGet)
	router.HandleFunc("/blogs", h.GetAllBlogs).Methods(http.MethodGet)
	router.HandleFunc("/blogs/author/{authorID}", h.GetBlogsByAuthorID).Methods(http.MethodGet)
	router.HandleFunc("/blogs/{id}", h.UpdateBlog).Methods(http.MethodPut)
	router.HandleFunc("/blogs/{id}", h.DeleteBlog).Methods(http.MethodDelete)
}

func (h *BlogHandler) CreateBlog(w http.ResponseWriter, r *http.Request) {
	var blog models.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	blog.ID = uuid.New().String()

	createdBlog, err := h.blogService.CreateBlog(blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBlog)
}

func (h *BlogHandler) GetBlogByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	blog, err := h.blogService.GetBlogByID(id)
	if err != nil {
		http.Error(w, "blog not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(blog)
}

func (h *BlogHandler) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := h.blogService.GetAllBlogs()
	if err != nil {
		http.Error(w, "could not retrieve blogs", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(blogs)
}

func (h *BlogHandler) GetBlogsByAuthorID(w http.ResponseWriter, r *http.Request) {
	authorID := mux.Vars(r)["authorID"]
	blogs, err := h.blogService.GetBlogsByAuthor(authorID)
	if err != nil {
		http.Error(w, "could not retrieve blogs", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(blogs)
}

func (h *BlogHandler) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var blog models.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	updatedBlog, err := h.blogService.UpdateBlog(id, blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBlog)
}

func (h *BlogHandler) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.blogService.DeleteBlog(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
