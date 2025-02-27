package db

import (
	"database/sql"
	"super-shiharai-kun/internal/models"
	"time"
)

type InvoiceDB struct {
	db *sql.DB
}

func NewInvoiceDB(db *sql.DB) *InvoiceDB {
	return &InvoiceDB{db: db}
}

func (r *InvoiceDB) Create(invoice *models.Invoice) error {
	query := `
        INSERT INTO invoices (company_id, client_id, issue_date, payment_amount, fee, fee_rate, tax, tax_rate, total_amount, due_date, status)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
	_, err := r.db.Exec(query,
		invoice.CompanyID,
		invoice.ClientID,
		invoice.IssueDate,
		invoice.PaymentAmount,
		invoice.Fee,
		invoice.FeeRate,
		invoice.Tax,
		invoice.TaxRate,
		invoice.TotalAmount,
		invoice.DueDate,
		invoice.Status,
	)
	return err
}

func (r *InvoiceDB) FindByPeriod(start, end time.Time) ([]models.Invoice, error) {
	query := `SELECT * FROM invoices WHERE due_date BETWEEN ? AND ?`
	rows, err := r.db.Query(query, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []models.Invoice
	for rows.Next() {
		var invoice models.Invoice
		if err := rows.Scan(
			&invoice.ID,
			&invoice.CompanyID,
			&invoice.ClientID,
			&invoice.IssueDate,
			&invoice.PaymentAmount,
			&invoice.Fee,
			&invoice.FeeRate,
			&invoice.Tax,
			&invoice.TaxRate,
			&invoice.TotalAmount,
			&invoice.DueDate,
			&invoice.Status,
		); err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}
	return invoices, nil
}
