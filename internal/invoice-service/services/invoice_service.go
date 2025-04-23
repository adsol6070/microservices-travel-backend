package services

import (
	"context"
	"errors"
	"microservices-travel-backend/internal/invoice-service/domain/models"
	"microservices-travel-backend/internal/invoice-service/domain/ports"
	"time"

	"github.com/google/uuid"
)

type InvoiceService struct {
	invoiceRepo ports.InvoiceRepositoryPort
}

func NewInvoiceService(repo ports.InvoiceRepositoryPort) *InvoiceService {
	return &InvoiceService{invoiceRepo: repo}
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, req models.Invoice) (*models.Invoice, error) {
	invoice := models.Invoice{
		ID:             uuid.New(),
		UserID:         req.UserID,
		BookingID:      req.BookingID,
		InvoiceNumber:  req.InvoiceNumber,
		Currency:       req.Currency,
		Amount:         req.Amount,
		TaxAmount:      req.TaxAmount,
		DiscountAmount: req.DiscountAmount,
		FinalAmount:    req.FinalAmount,
		Status:         req.Status,
		PaymentMethod:  req.PaymentMethod,
		PaidAt:         req.PaidAt,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Description:    req.Description,
	}

	if s.invoiceRepo.InvoiceNumberExists(ctx, req.InvoiceNumber) {
		return nil, errors.New("invoice number already exists")
	}

	createdInvoice, err := s.invoiceRepo.Create(ctx, &invoice)
	if err != nil {
		return nil, errors.New("failed to create invoice")
	}
	return createdInvoice, nil
}

func (s *InvoiceService) GetInvoiceByID(ctx context.Context, id string) (*models.Invoice, error) {
	invoice, err := s.invoiceRepo.GetByID(ctx, id)
	if err != nil || invoice == nil {
		return nil, errors.New("invoice not found")
	}
	return invoice, nil
}

func (s *InvoiceService) GetInvoicesByUserID(ctx context.Context, userID string) ([]*models.Invoice, error) {
	invoices, err := s.invoiceRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, errors.New("failed to retrieve invoices by user")
	}
	return invoices, nil
}

func (s *InvoiceService) GetInvoiceByBookingID(ctx context.Context, bookingID string) (*models.Invoice, error) {
	invoice, err := s.invoiceRepo.GetByBookingID(ctx, bookingID)
	if err != nil || invoice == nil {
		return nil, errors.New("invoice not found for booking")
	}
	return invoice, nil
}

func (s *InvoiceService) GetAllInvoices(ctx context.Context) ([]*models.Invoice, error) {
	invoices, err := s.invoiceRepo.GetAll(ctx)
	if err != nil {
		return nil, errors.New("failed to retrieve invoices")
	}
	return invoices, nil
}

func (s *InvoiceService) UpdateInvoice(ctx context.Context, id string, req models.UpdateInvoiceRequest) (*models.Invoice, error) {
	existingInvoice, err := s.invoiceRepo.GetByID(ctx, id)
	if err != nil || existingInvoice == nil {
		return nil, errors.New("invoice not found")
	}

	if req.Status != nil {
		existingInvoice.Status = *req.Status
	}
	if req.PaymentMethod != nil {
		existingInvoice.PaymentMethod = *req.PaymentMethod
	}
	if req.PaidAt != nil {
		existingInvoice.PaidAt = req.PaidAt
	}

	existingInvoice.UpdatedAt = time.Now()

	updatedInvoice, err := s.invoiceRepo.Update(ctx, id, existingInvoice)
	if err != nil {
		return nil, errors.New("failed to update invoice")
	}
	return updatedInvoice, nil
}

func (s *InvoiceService) DeleteInvoice(ctx context.Context, id string) error {
	invoice, err := s.invoiceRepo.GetByID(ctx, id)
	if err != nil || invoice == nil {
		return errors.New("invoice not found")
	}

	if err := s.invoiceRepo.Delete(ctx, id); err != nil {
		return errors.New("failed to delete invoice")
	}
	return nil
}
