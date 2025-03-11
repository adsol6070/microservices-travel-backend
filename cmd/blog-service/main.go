package main

import (
	"log"
	"microservices-travel-backend/internal/blog-service/adapters/handlers"
	"microservices-travel-backend/internal/blog-service/adapters/repositories"
	"microservices-travel-backend/internal/blog-service/services"
	middleware "microservices-travel-backend/pkg/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	blogRepo, err := repositories.NewPostgreSQLBlogRepository()
	if err != nil {
		log.Fatalf("Failed to create blog repository: %v", err)
	}
	blogCatgeoryRepo, err := repositories.NewPostgreSQLBlogCategoryRepository()
	if err != nil {
		log.Fatalf("Failed to create blog category repository: %v", err)
	}

	blogService := services.NewBlogService(blogRepo)
	blogCategoryService := services.NewBlogCategoryService(blogCatgeoryRepo)

	blogHandler := handlers.NewBlogHandler(blogService)
	blogCategoryHandler := handlers.NewBlogCategoryHandler(blogCategoryService)

	router := mux.NewRouter()
	blogHandler.RegisterRoutes(router)
	blogCategoryHandler.RegisterRoutes(router)

	port := ":7200"
	log.Printf("Starting blog service on port %s...", port)
	if err := http.ListenAndServe(port, middleware.CORSMiddleware(router)); err != nil {
		log.Fatalf("Failed to start blog service: %v", err)
	}
}
