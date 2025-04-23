package main

import (
	"log"
	"microservices-travel-backend/internal/tour-and-packages/adapters/handlers"
	"microservices-travel-backend/internal/tour-and-packages/adapters/repositories"
	"microservices-travel-backend/internal/tour-and-packages/services"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "tour-and-packages-service: ", log.LstdFlags|log.Lshortfile)

	packageRepo, err := repositories.NewPackageRepository()
	tourRepo, err := repositories.NewTourRepository()

	packageService := services.NewPackageService(packageRepo)
	tourService := services.NewTourService(tourRepo)

	packageHandler := handlers.NewPackageHandler(packageService)
	tourHandler := handlers.NewTourHandler(tourService)

	router := mux.NewRouter()
	packageHandler.RegisterRoutes(router)
	tourHandler.RegisterRoutes(router)

	port := "7200"
	logger.Printf("Starting server on port %s...\n", "7200")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
