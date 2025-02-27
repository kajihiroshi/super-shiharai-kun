package service

import (
	"super-shiharai-kun/internal/models"
	"testing"
	"time"
)

// MockInvoiceDB simulates the database interactions for testing purposes.
type MockInvoiceDB struct {
	CreateFunc       func(invoice *models.Invoice) error
	FindByPeriodFunc func(start, end time.Time) ([]models.Invoice, error)
}

// Create is the mock implementation of the Create method of the InvoiceDBInterface.
func (m *MockInvoiceDB) Create(invoice *models.Invoice) error {
	return m.CreateFunc(invoice)
}

// FindByPeriod is the mock implementation of the FindByPeriod method of the InvoiceDBInterface.
func (m *MockInvoiceDB) FindByPeriod(start, end time.Time) ([]models.Invoice, error) {
	return m.FindByPeriodFunc(start, end)
}

func TestInvoiceService_CreateInvoice(t *testing.T) {
	mockDB := &MockInvoiceDB{
		CreateFunc: func(invoice *models.Invoice) error {
			invoice.ID = 1
			return nil
		},
	}

	// Now using db.InvoiceDBInterface (not the concrete *db.InvoiceDB)
	invoiceService := NewInvoiceService(mockDB)

	invoice := &models.Invoice{
		CompanyID:     1,
		ClientID:      1,
		PaymentAmount: 10000,
		DueDate:       time.Now().AddDate(0, 1, 0),
	}

	err := invoiceService.CreateInvoice(invoice)
	if err != nil {
		t.Fatalf("Failed to create invoice: %v", err)
	}

	if invoice.ID != 1 {
		t.Error("Expected invoice ID to be 1, got", invoice.ID)
	}

	if invoice.TotalAmount != 10440 {
		t.Error("Expected total amount to be 10440, got", invoice.TotalAmount)
	}
}

func TestInvoiceService_GetInvoicesByPeriod(t *testing.T) {
	mockDB := &MockInvoiceDB{
		FindByPeriodFunc: func(start, end time.Time) ([]models.Invoice, error) {
			return []models.Invoice{
				{
					ID:            1,
					CompanyID:     1,
					ClientID:      1,
					PaymentAmount: 10000,
					TotalAmount:   10440,
					DueDate:       time.Now(),
				},
			}, nil
		},
	}

	// Using db.InvoiceDBInterface (mockDB satisfies this interface)
	invoiceService := NewInvoiceService(mockDB)

	start := time.Now().AddDate(0, -1, 0)
	end := time.Now().AddDate(0, 1, 0)

	invoices, err := invoiceService.GetInvoicesByPeriod(start, end)
	if err != nil {
		t.Fatalf("Failed to get invoices: %v", err)
	}

	if len(invoices) != 1 {
		t.Error("Expected 1 invoice, got", len(invoices))
	}
}
