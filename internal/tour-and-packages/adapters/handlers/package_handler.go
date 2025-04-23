package handlers

import (
	"microservices-travel-backend/internal/tour-and-packages/domain/ports"
	"net/http"

	"github.com/gorilla/mux"
)

type PackageHandler struct {
	packageService ports.PackageServicePort
}

func NewPackageHandler(service ports.PackageServicePort) *PackageHandler {
	return &PackageHandler{packageService: service}
}

func (h *PackageHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/packages", h.CreatePackage).Methods("POST")
	router.HandleFunc("/packages/{id}", h.GetPackageByID).Methods("GET")
	router.HandleFunc("/packages", h.GetAllPackages).Methods("GET")
	router.HandleFunc("/packages/{id}", h.UpdatePackage).Methods("PATCH")
	router.HandleFunc("/packages/{id}", h.DeletePackage).Methods("DELETE")
}

func (h *PackageHandler) CreatePackage(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a tour
}

func (h *PackageHandler) GetPackageByID(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a tour by ID
}

func (h *PackageHandler) GetAllPackages(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting all tours
}

func (h *PackageHandler) UpdatePackage(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a tour
}

func (h *PackageHandler) DeletePackage(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a tour
}
