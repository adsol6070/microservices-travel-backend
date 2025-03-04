package handlers

import (
	"encoding/json"
	"net/http"

	"microservices-travel-backend/internal/blog-service/domain/models"
	"microservices-travel-backend/internal/blog-service/domain/ports"
	"microservices-travel-backend/pkg/response"
	validator "microservices-travel-backend/pkg/validation"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
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
		response.BadRequest(w, "Invalid request body")
		return
	}
	if err := validator.ValidateStruct(blog); err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	blog.ID = uuid.New().String()
	blog.Tags = pq.StringArray(blog.Tags)

	createdBlog, err := h.blogService.CreateBlog(r.Context(), &blog)
	if err != nil {
		response.InternalServerError(w, "Failed to create blog")
		return
	}

	response.Success(w, http.StatusCreated, "Blog created successfully", createdBlog)
}
func (h *BlogHandler) GetBlogByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	blog, err := h.blogService.GetBlogByID(r.Context(), id)
	if err != nil {
		response.NotFound(w, "Blog not found")
		return
	}

	response.Success(w, http.StatusOK, "Blog retrieved successfully", blog)
}

func (h *BlogHandler) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := h.blogService.GetAllBlogs(r.Context())
	if err != nil {
		response.InternalServerError(w, "Could not retrieve blogs")
		return
	}

	response.Success(w, http.StatusOK, "Blogs retrieved successfully", blogs)
}

func (h *BlogHandler) GetBlogsByAuthorID(w http.ResponseWriter, r *http.Request) {
	authorID := mux.Vars(r)["authorID"]

	blogs, err := h.blogService.GetBlogsByAuthor(r.Context(), authorID)
	if err != nil {
		response.InternalServerError(w, "Could not retrieve blogs")
		return
	}

	response.Success(w, http.StatusOK, "Blogs retrieved successfully", blogs)
}

func (h *BlogHandler) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var blog models.Blog

	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		response.BadRequest(w, "Invalid request payload")
		return
	}

	if err := validator.ValidateStruct(blog); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	updatedBlog, err := h.blogService.UpdateBlog(r.Context(), id, &blog)
	if err != nil {
		response.InternalServerError(w, "Failed to update blog")
		return
	}

	response.Success(w, http.StatusOK, "Blog updated successfully", updatedBlog)
}

func (h *BlogHandler) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.blogService.DeleteBlog(r.Context(), id); err != nil {
		response.InternalServerError(w, "Failed to delete blog")
		return
	}

	response.Success(w, http.StatusNoContent, "Blog deleted successfully", nil)
}
