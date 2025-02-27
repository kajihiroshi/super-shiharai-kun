package models

import "time"

type Invoice struct {
	ID            int
	CompanyID     int
	ClientID      int
	IssueDate     time.Time
	PaymentAmount float64
	Fee           float64
	FeeRate       float64
	Tax           float64
	TaxRate       float64
	TotalAmount   float64
	DueDate       time.Time
	Status        string // e.g., "未処理", "処理中", "支払い済み", "エラー"
}

// CalculateTotalAmount calculates Fee, Tax, and TotalAmount for the invoice
func (i *Invoice) CalculateTotalAmount() {
	// Calculate the fee and tax based on the rates and payment amount
	i.Fee = i.PaymentAmount * i.FeeRate
	i.Tax = i.PaymentAmount * i.TaxRate
	// Calculate the total amount
	i.TotalAmount = i.PaymentAmount + i.Fee + i.Tax
}
