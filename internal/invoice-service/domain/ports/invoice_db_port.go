package ports

import (
	"context"
	"microservices-travel-backend/internal/invoice-service/domain/models"
)

type InvoiceRepositoryPort interface {
	Create(ctx context.Context, invoice *models.Invoice) (*models.Invoice, error)
	GetByID(ctx context.Context, id string) (*models.Invoice, error)
	GetByUserID(ctx context.Context, userID string) ([]*models.Invoice, error)
	GetByBookingID(ctx context.Context, bookingID string) (*models.Invoice, error)
	GetAll(ctx context.Context) ([]*models.Invoice, error)
	Update(ctx context.Context, id string, updatedInvoice *models.Invoice) (*models.Invoice, error)
	Delete(ctx context.Context, id string) error
	InvoiceNumberExists(ctx context.Context, invoiceNumber string) bool
}
