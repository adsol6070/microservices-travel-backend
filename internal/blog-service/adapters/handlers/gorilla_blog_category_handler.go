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
)

type BlogCategoryHandler struct {
	categoryService ports.BlogCategoryServicePort
}

func NewBlogCategoryHandler(service ports.BlogCategoryServicePort) *BlogCategoryHandler {
	return &BlogCategoryHandler{categoryService: service}
}

func (h *BlogCategoryHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/blogCategory", h.CreateCategory).Methods(http.MethodPost)
	router.HandleFunc("/blogCategory/{id}", h.GetCategoryByID).Methods(http.MethodGet)
	router.HandleFunc("/blogCategories", h.GetAllCategories).Methods(http.MethodGet)
	router.HandleFunc("/blogCategory/{id}", h.UpdateCategory).Methods(http.MethodPut)
	router.HandleFunc("/blogCategory/{id}", h.DeleteCategory).Methods(http.MethodDelete)
}

func (h *BlogCategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.BlogCategory
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}
	if err := validator.ValidateStruct(category); err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	category.ID = uuid.New().String()

	createdCategory, err := h.categoryService.CreateCategory(r.Context(), &category)
	if err != nil {
		response.InternalServerError(w, "Failed to create category")
		return
	}

	response.Success(w, http.StatusCreated, "Category created successfully", createdCategory)
}

func (h *BlogCategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	category, err := h.categoryService.GetCategoryByID(r.Context(), id)
	if err != nil {
		response.NotFound(w, "Category not found")
		return
	}

	response.Success(w, http.StatusOK, "Category retrieved successfully", category)
}

func (h *BlogCategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoryService.GetAllCategories(r.Context())
	if err != nil {
		response.InternalServerError(w, "Could not retrieve categories")
		return
	}

	response.Success(w, http.StatusOK, "Categories retrieved successfully", categories)
}

func (h *BlogCategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var category models.BlogCategory

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		response.BadRequest(w, "Invalid request payload")
		return
	}

	if err := validator.ValidateStruct(category); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	updatedCategory, err := h.categoryService.UpdateCategory(r.Context(), id, &category)
	if err != nil {
		response.InternalServerError(w, "Failed to update category")
		return
	}

	response.Success(w, http.StatusOK, "Category updated successfully", updatedCategory)
}

func (h *BlogCategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.categoryService.DeleteCategory(r.Context(), id); err != nil {
		response.InternalServerError(w, "Failed to delete category")
		return
	}

	response.Success(w, http.StatusNoContent, "Category deleted successfully", nil)
}
