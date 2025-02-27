package models

import (
	"testing"
)

func TestInvoice_CalculateTotalAmount(t *testing.T) {
	invoice := Invoice{
		PaymentAmount: 10000,
		FeeRate:       0.04,
		TaxRate:       0.10,
	}

	// Call the method to calculate the total amount
	invoice.CalculateTotalAmount()

	// Check if the fee is correct
	if invoice.Fee != 400 {
		t.Errorf("Expected fee to be 400, got %f", invoice.Fee)
	}

	// Check if the tax is correct
	if invoice.Tax != 1000 {
		t.Errorf("Expected tax to be 1000, got %f", invoice.Tax)
	}

	// Check if the total amount is correct
	if invoice.TotalAmount != 11400 {
		t.Errorf("Expected total amount to be 11400, got %f", invoice.TotalAmount)
	}
}
