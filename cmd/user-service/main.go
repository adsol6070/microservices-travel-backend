package main

import (
	"log"
	"microservices-travel-backend/internal/user-service/adapters/handlers"
	"microservices-travel-backend/internal/user-service/adapters/repositories"
	"microservices-travel-backend/internal/user-service/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	userRepo, err := repositories.NewPostgreSQLUserRepository()

	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router := mux.NewRouter()
	userHandler.RegisterRoutes(router)

	port := ":5001"
	log.Printf("Starting user service on port %s...", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
