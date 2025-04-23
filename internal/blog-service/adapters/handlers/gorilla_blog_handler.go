package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"microservices-travel-backend/internal/blog-service/domain/models"
	"microservices-travel-backend/internal/blog-service/domain/ports"
	middleware "microservices-travel-backend/pkg/middlewares"
	"microservices-travel-backend/pkg/response"
	validator "microservices-travel-backend/pkg/validation"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type BlogHandler struct {
	blogService ports.BlogServicePort
}

func NewBlogHandler(service ports.BlogServicePort) *BlogHandler {
	return &BlogHandler{blogService: service}
}

func (h *BlogHandler) RegisterRoutes(router *mux.Router, imageUploadMiddleware middleware.ImageUploadMiddleware) {
	router.HandleFunc("/blogs", func(w http.ResponseWriter, r *http.Request) {
		imageUploadMiddleware.UploadImage("blog-service", "blogImage", h.CreateBlog).ServeHTTP(w, r)
	}).Methods(http.MethodPost)
	router.HandleFunc("/blogs/{id}", h.GetBlogByID).Methods(http.MethodGet)
	router.HandleFunc("/blogs/category/{category}", h.GetBlogByCategory).Methods(http.MethodGet)
	router.HandleFunc("/blogs", h.GetAllBlogs).Methods(http.MethodGet)
	router.HandleFunc("/blogs/author/{authorID}", h.GetBlogsByAuthorID).Methods(http.MethodGet)
	router.HandleFunc("/blogs/{id}", func(w http.ResponseWriter, r *http.Request) {
		imageUploadMiddleware.UploadImage("blog-service", "blogImage", h.UpdateBlog).ServeHTTP(w, r)
	}).Methods(http.MethodPatch)
	router.HandleFunc("/blogs/{id}", h.DeleteBlog).Methods(http.MethodDelete)
	router.HandleFunc("/blogs/{id}/status", h.UpdateBlogStatus).Methods(http.MethodPatch)
}

func (h *BlogHandler) CreateBlog(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Printf("Error parsing form data: %v", err)
		response.BadRequest(w, "Invalid form data: "+err.Error())
		return
	}

	categoryIDStr := r.FormValue("category_id")
	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		log.Printf("Invalid category_id: %v", err)
		response.BadRequest(w, "Invalid category_id format")
		return
	}

	var publishedAt *time.Time
	publishedAtStr := r.FormValue("published_at")
	if publishedAtStr != "" {
		parsedTime, err := time.Parse(time.RFC3339, publishedAtStr)
		if err != nil {
			log.Printf("Invalid published_at: %v", err)
			response.BadRequest(w, "Invalid published_at format (expected YYYY-MM-DD HH:MM:SS)")
			return
		}
		publishedAt = &parsedTime
	}

	imageURLs, ok := middleware.GetImageURLsFromContext(r.Context())
	if !ok || len(imageURLs) == 0 {
		log.Println("No image uploaded")
		response.BadRequest(w, "Blog image is required")
		return
	}

	blogRequest := models.CreateBlogRequest{
		Title:           r.FormValue("title"),
		Slug:            r.FormValue("slug"),
		Author:          r.FormValue("author"),
		CategoryID:      categoryID,
		Content:         r.FormValue("content"),
		Excerpt:         r.FormValue("excerpt"),
		MetaTitle:       r.FormValue("meta_title"),
		MetaDescription: r.FormValue("meta_description"),
		Status:          r.FormValue("status"),
		Thumbnail:       imageURLs[0],
		Tags:            nil,
		PublishedAt:     publishedAt,
	}

	tagsJSON := r.FormValue("tags")

	fmt.Println("Received tagsJSON:", tagsJSON)
	if tagsJSON != "" {
		var tags []string
		err := json.Unmarshal([]byte(tagsJSON), &tags)
		if err != nil {
			http.Error(w, "Invalid tags format", http.StatusBadRequest)
			return
		}
		fmt.Println("Parsed tags array:", tags)
		blogRequest.Tags = tags
	}

	if err := validator.ValidateStruct(blogRequest); err != nil {
		log.Printf("Validation error: %v", err)
		response.BadRequest(w, err.Error())
		return
	}

	_, err = h.blogService.CreateBlog(r.Context(), blogRequest)
	if err != nil {
		log.Printf("Error creating blog: %v", err)
		response.InternalServerError(w, "Failed to create blog")
		return
	}

	response.Success(w, http.StatusCreated, "Blog created successfully", nil)
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

func (h *BlogHandler) GetBlogByCategory(w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]

	blog, err := h.blogService.GetBlogByCategory(r.Context(), category)
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
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Printf("Error parsing form data: %v", err)
		response.BadRequest(w, "Invalid form data: "+err.Error())
		return
	}

	categoryIDStr := r.FormValue("category_id")
	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		log.Printf("Invalid category_id: %v", err)
		response.BadRequest(w, "Invalid category_id format")
		return
	}

	var publishedAt *time.Time
	publishedAtStr := r.FormValue("published_at")
	if publishedAtStr != "" {
		parsedTime, err := time.Parse(time.RFC3339, publishedAtStr)
		if err != nil {
			log.Printf("Invalid published_at: %v", err)
			response.BadRequest(w, "Invalid published_at format (expected YYYY-MM-DD HH:MM:SS)")
			return
		}
		publishedAt = &parsedTime
	}

	imageURLs, ok := middleware.GetImageURLsFromContext(r.Context())
	if !ok || len(imageURLs) == 0 {
		log.Println("No new image uploaded, keeping the old one")
	}

	blogUpdate := models.UpdateBlogRequest{
		Title:           toPtr(r.FormValue("title")),
		Slug:            toPtr(r.FormValue("slug")),
		Author:          toPtr(r.FormValue("author")),
		CategoryID:      &categoryID,
		Content:         toPtr(r.FormValue("content")),
		Excerpt:         toPtr(r.FormValue("excerpt")),
		MetaTitle:       toPtr(r.FormValue("meta_title")),
		MetaDescription: toPtr(r.FormValue("meta_description")),
		Status:          toPtr(r.FormValue("status")),
		PublishedAt:     publishedAt,
	}

	if len(imageURLs) > 0 {
		blogUpdate.Thumbnail = toPtr(imageURLs[0])
	}

	tagsJSON := r.FormValue("tags")
	fmt.Println("Received tagsJSON:", tagsJSON)

	if tagsJSON != "" {
		var tags []string
		err := json.Unmarshal([]byte(tagsJSON), &tags)
		if err != nil {
			response.BadRequest(w, "Invalid tags format")
			return
		}
		blogUpdate.Tags = tags
	}

	if err := validator.ValidateStruct(blogUpdate); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	updatedBlog, err := h.blogService.UpdateBlog(r.Context(), id, blogUpdate)
	if err != nil {
		response.InternalServerError(w, "Failed to update blog")
		return
	}

	response.Success(w, http.StatusOK, "Blog updated successfully", updatedBlog)
}

func toPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func (h *BlogHandler) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.blogService.DeleteBlog(r.Context(), id); err != nil {
		response.InternalServerError(w, "Failed to delete blog")
		return
	}

	response.Success(w, http.StatusNoContent, "Blog deleted successfully", nil)
}

func (h *BlogHandler) UpdateBlogStatus(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var statusUpdate struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&statusUpdate); err != nil {
		response.BadRequest(w, "Invalid request payload")
		return
	}

	log.Printf("Received request to update blog ID: %s with status: %s", id, statusUpdate.Status)

	updatedBlog, err := h.blogService.UpdateBlogStatus(r.Context(), id, statusUpdate.Status)
	if err != nil {
		response.InternalServerError(w, "Failed to update blog status")
		return
	}

	response.Success(w, http.StatusOK, "Blog status updated successfully", updatedBlog)
}
