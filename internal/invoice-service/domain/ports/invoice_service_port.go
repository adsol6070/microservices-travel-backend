package ports

import (
	"context"
	"microservices-travel-backend/internal/invoice-service/domain/models"
)

type InvoiceServicePort interface {
	CreateInvoice(ctx context.Context, invoice models.Invoice) (*models.Invoice, error)
	GetInvoiceByID(ctx context.Context, id string) (*models.Invoice, error)
	GetInvoicesByUserID(ctx context.Context, userID string) ([]*models.Invoice, error)
	GetInvoiceByBookingID(ctx context.Context, bookingID string) (*models.Invoice, error)
	GetAllInvoices(ctx context.Context) ([]*models.Invoice, error)
	UpdateInvoice(ctx context.Context, id string, invoice models.UpdateInvoiceRequest) (*models.Invoice, error)
	DeleteInvoice(ctx context.Context, id string) error
}
