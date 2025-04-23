package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"microservices-travel-backend/internal/invoice-service/domain/models"
	"microservices-travel-backend/internal/invoice-service/domain/ports"
	"microservices-travel-backend/pkg/response"
	validator "microservices-travel-backend/pkg/validation"

	"github.com/gorilla/mux"
)

type InvoiceHandler struct {
	invoiceService ports.InvoiceServicePort
}

func NewInvoiceHandler(service ports.InvoiceServicePort) *InvoiceHandler {
	return &InvoiceHandler{invoiceService: service}
}

func (h *InvoiceHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/invoices", h.CreateInvoice).Methods(http.MethodPost)
	router.HandleFunc("/invoices", h.GetAllInvoices).Methods(http.MethodGet)
	router.HandleFunc("/invoices/{id}", h.GetInvoiceByID).Methods(http.MethodGet)
	router.HandleFunc("/invoices/user/{userID}", h.GetInvoicesByUserID).Methods(http.MethodGet)
	router.HandleFunc("/invoices/booking/{bookingID}", h.GetInvoiceByBookingID).Methods(http.MethodGet)
	router.HandleFunc("/invoices/{id}", h.UpdateInvoice).Methods(http.MethodPut)
	router.HandleFunc("/invoices/{id}", h.DeleteInvoice).Methods(http.MethodDelete)
}

func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var req models.Invoice
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}
	if err := validator.ValidateStruct(req); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	invoice, err := h.invoiceService.CreateInvoice(r.Context(), req)
	if err != nil {
		response.InternalServerError(w, "Failed to create invoice")
		return
	}

	response.Success(w, http.StatusCreated, "Invoice created successfully", invoice)
}

func (h *InvoiceHandler) GetInvoiceByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	invoice, err := h.invoiceService.GetInvoiceByID(r.Context(), id)
	if err != nil {
		response.NotFound(w, "Invoice not found")
		return
	}

	response.Success(w, http.StatusOK, "Invoice retrieved successfully", invoice)
}

func (h *InvoiceHandler) GetAllInvoices(w http.ResponseWriter, r *http.Request) {
	invoices, err := h.invoiceService.GetAllInvoices(r.Context())
	if err != nil {
		log.Println("Error retrieving invoices:", err)
		response.InternalServerError(w, "Failed to retrieve invoices")
		return
	}

	response.Success(w, http.StatusOK, "Invoices retrieved successfully", invoices)
}

func (h *InvoiceHandler) GetInvoicesByUserID(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]

	invoices, err := h.invoiceService.GetInvoicesByUserID(r.Context(), userID)
	if err != nil {
		response.InternalServerError(w, "Failed to retrieve invoices")
		return
	}

	response.Success(w, http.StatusOK, "Invoices retrieved successfully", invoices)
}

func (h *InvoiceHandler) GetInvoiceByBookingID(w http.ResponseWriter, r *http.Request) {
	bookingID := mux.Vars(r)["bookingID"]

	invoice, err := h.invoiceService.GetInvoiceByBookingID(r.Context(), bookingID)
	if err != nil {
		response.NotFound(w, "Invoice not found")
		return
	}

	response.Success(w, http.StatusOK, "Invoice retrieved successfully", invoice)
}

func (h *InvoiceHandler) UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var req models.UpdateInvoiceRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request payload")
		return
	}
	if err := validator.ValidateStruct(req); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	invoice, err := h.invoiceService.UpdateInvoice(r.Context(), id, req)
	if err != nil {
		response.InternalServerError(w, "Failed to update invoice")
		return
	}

	response.Success(w, http.StatusOK, "Invoice updated successfully", invoice)
}

func (h *InvoiceHandler) DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.invoiceService.DeleteInvoice(r.Context(), id); err != nil {
		response.InternalServerError(w, "Failed to delete invoice")
		return
	}

	response.Success(w, http.StatusNoContent, "Invoice deleted successfully", nil)
}
