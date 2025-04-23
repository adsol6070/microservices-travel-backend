package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"microservices-travel-backend/internal/invoice-service/domain/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQLInvoiceRepository struct {
	db *gorm.DB
}

func NewPostgreSQLInvoiceRepository() (*PostgreSQLInvoiceRepository, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")
	sslMode := os.Getenv("DATABASE_SSLMODE")

	var dsn string
	if databaseName == "" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/?sslmode=%s",
			databaseUsername, databasePassword, databaseURL, databasePort, sslMode)
	} else {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			databaseUsername, databasePassword, databaseURL, databasePort, databaseName, sslMode)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to the database")

	return &PostgreSQLInvoiceRepository{db: db}, nil
}

func (r *PostgreSQLInvoiceRepository) Create(ctx context.Context, invoice *models.Invoice) (*models.Invoice, error) {
	if err := r.db.WithContext(ctx).Table("invoices").Create(invoice).Error; err != nil {
		return nil, err
	}
	return invoice, nil
}

func (r *PostgreSQLInvoiceRepository) GetByID(ctx context.Context, id string) (*models.Invoice, error) {
	var invoice models.Invoice
	if err := r.db.WithContext(ctx).Table("invoices").Preload("Description").First(&invoice, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &invoice, nil
}

func (r *PostgreSQLInvoiceRepository) GetByUserID(ctx context.Context, userID string) ([]*models.Invoice, error) {
	var invoices []*models.Invoice
	if err := r.db.WithContext(ctx).Table("invoices").Preload("Description").Where("user_id = ?", userID).Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}

func (r *PostgreSQLInvoiceRepository) GetByBookingID(ctx context.Context, bookingID string) (*models.Invoice, error) {
	var invoice models.Invoice
	if err := r.db.WithContext(ctx).Table("invoices").Preload("Description").Where("booking_id = ?", bookingID).First(&invoice).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &invoice, nil
}

func (r *PostgreSQLInvoiceRepository) GetAll(ctx context.Context) ([]*models.Invoice, error) {
	var invoices []*models.Invoice
	if err := r.db.WithContext(ctx).Table("invoices").Preload("Description").Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}

func (r *PostgreSQLInvoiceRepository) Update(ctx context.Context, id string, invoice *models.Invoice) (*models.Invoice, error) {
	if err := r.db.WithContext(ctx).Table("invoices").Save(invoice).Error; err != nil {
		return nil, err
	}
	return invoice, nil
}

func (r *PostgreSQLInvoiceRepository) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Table("invoices").Delete(&models.Invoice{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgreSQLInvoiceRepository) InvoiceNumberExists(ctx context.Context, invoiceNumber string) bool {
	var count int64
	err := r.db.WithContext(ctx).Table("invoices").Where("invoice_number = ?", invoiceNumber).Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}
