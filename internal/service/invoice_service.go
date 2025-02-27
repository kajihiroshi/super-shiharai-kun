package service

import (
	"super-shiharai-kun/internal/db"
	"super-shiharai-kun/internal/models"
	"time"
)

type InvoiceService struct {
	db db.InvoiceDBInterface
}

func NewInvoiceService(db db.InvoiceDBInterface) *InvoiceService {
	return &InvoiceService{db: db}
}

func (s *InvoiceService) CreateInvoice(invoice *models.Invoice) error {
	// Calculate total amount with fee and tax
	invoice.FeeRate = 0.04
	invoice.TaxRate = 0.10
	invoice.Fee = invoice.PaymentAmount * invoice.FeeRate
	invoice.Tax = invoice.Fee * invoice.TaxRate
	invoice.TotalAmount = invoice.PaymentAmount + invoice.Fee + invoice.Tax

	return s.db.Create(invoice)
}

func (s *InvoiceService) GetInvoicesByPeriod(start, end time.Time) ([]models.Invoice, error) {
	return s.db.FindByPeriod(start, end)
}
