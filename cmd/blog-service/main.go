package main

import (
	"log"
	"microservices-travel-backend/internal/blog-service/adapters/handlers"
	"microservices-travel-backend/internal/blog-service/adapters/repositories"
	"microservices-travel-backend/internal/blog-service/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	blogRepo, err := repositories.NewPostgreSQLBlogRepository()
	if err != nil {
		log.Fatalf("Failed to create blog repository: %v", err)
	}

	blogService := services.NewBlogService(blogRepo)

	blogHandler := handlers.NewBlogHandler(blogService)

	router := mux.NewRouter()
	blogHandler.RegisterRoutes(router)

	port := ":7200" 
	log.Printf("Starting blog service on port %s...", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start blog service: %v", err)
	}
}
