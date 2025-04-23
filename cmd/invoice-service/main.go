package main

import (
	"log"
	"net/http"

	"microservices-travel-backend/internal/invoice-service/adapters/handlers"
	"microservices-travel-backend/internal/invoice-service/adapters/repositories"
	"microservices-travel-backend/internal/invoice-service/services"
	middleware "microservices-travel-backend/pkg/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	// Init PostgreSQL invoice repository
	invoiceRepo, err := repositories.NewPostgreSQLInvoiceRepository()
	if err != nil {
		log.Fatalf("Failed to initialize invoice repository: %v", err)
	}

	// Initialize service
	invoiceService := services.NewInvoiceService(invoiceRepo)

	// Initialize handler
	invoiceHandler := handlers.NewInvoiceHandler(invoiceService)

	// Setup router and register routes
	router := mux.NewRouter()
	invoiceHandler.RegisterRoutes(router)

	// Start HTTP server
	port := ":8200"
	log.Printf("Starting invoice service on port %s...", port)
	if err := http.ListenAndServe(port, middleware.CORSMiddleware(router)); err != nil {
		log.Fatalf("Failed to start invoice service: %v", err)
	}
}
