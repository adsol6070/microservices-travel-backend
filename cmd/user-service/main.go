package main

import (
	"log"
	"microservices-travel-backend/internal/shared/rabbitmq"
	"microservices-travel-backend/internal/user-service/app/usecase"
	"microservices-travel-backend/internal/user-service/domain/auth"
	"microservices-travel-backend/internal/user-service/domain/email"
	"microservices-travel-backend/internal/user-service/domain/user"
	"microservices-travel-backend/internal/user-service/infrastructure/handlers"
	"microservices-travel-backend/internal/user-service/infrastructure/persistance/postgres"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	logger := log.New(os.Stdout, "user-service: ", log.LstdFlags|log.Lshortfile)

	authRepo, err := postgres.NewAuthRepository()
	if err != nil {
		logger.Fatalf("Failed to initialize user repository: %v", err)
	}

	userRepo, err := postgres.NewUserRepository()
	if err != nil {
		logger.Fatalf("Failed to initialize auth repository: %v", err)
	}

	emailService := email.NewEmailService(rabbitmq.New("emailQueue", os.Getenv("RABBITMQ_URL")))

	// Initialize domain services
	authService := auth.NewAuthService(userRepo, authRepo, logger)
	userService := user.NewUserService(userRepo)

	// Initialize use cases
	authUsecase := usecase.NewAuthUsecase(authService, emailService)
	userUsecase := usecase.NewUserUsecase(userService)

	router := mux.NewRouter()

	// Initialize API Handlers
	handlers.NewAuthHandler(router, authUsecase)
	handlers.NewUserHandler(router, userUsecase)

	// Set Port and Start Server
	serverPort := "7100"
	logger.Printf("Starting server on port %s...\n", "7100")
	if err := http.ListenAndServe(":"+serverPort, router); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}

}
