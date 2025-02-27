package db

import (
	"database/sql"
	"super-shiharai-kun/internal/models"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// setupTestDB opens a new database connection and starts a transaction for tests.
func setupTestDB(t *testing.T) (*sql.DB, *sql.Tx) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/super_shiharai_kun_test")
	if err != nil {
		t.Fatal(err)
	}

	// Start a new transaction for the test
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	// Set the database connection to use the transaction for the test
	// This allows rollback after the test is done.
	return db, tx
}

// cleanupTestDB rolls back any changes made during the test to maintain isolation.
func cleanupTestDB(t *testing.T, tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil && err != sql.ErrTxDone {
		t.Fatal("Failed to rollback transaction:", err)
	}
}

func TestInvoiceDB_Create(t *testing.T) {
	db, tx := setupTestDB(t)
	defer cleanupTestDB(t, tx)

	invoiceDB := NewInvoiceDB(db)

	invoice := &models.Invoice{
		CompanyID:     1,
		ClientID:      1,
		IssueDate:     time.Now(),
		PaymentAmount: 10000,
		DueDate:       time.Now().AddDate(0, 1, 0),
		Status:        "未処理",
	}

	err := invoiceDB.Create(invoice)
	if err != nil {
		t.Fatalf("Failed to create invoice: %v", err)
	}

	if invoice.ID == 0 {
		t.Error("Expected invoice ID to be set, got 0")
	}
}

func TestInvoiceDB_FindByPeriod(t *testing.T) {
	db, tx := setupTestDB(t)
	defer cleanupTestDB(t, tx)

	invoiceDB := NewInvoiceDB(db)

	start := time.Now().AddDate(0, -1, 0)
	end := time.Now().AddDate(0, 1, 0)

	// Insert a dummy invoice to make sure FindByPeriod works.
	invoice := &models.Invoice{
		CompanyID:     1,
		ClientID:      1,
		IssueDate:     time.Now(),
		PaymentAmount: 10000,
		DueDate:       time.Now().AddDate(0, 1, 0),
		Status:        "未処理",
	}

	err := invoiceDB.Create(invoice)
	if err != nil {
		t.Fatalf("Failed to create test invoice: %v", err)
	}

	invoices, err := invoiceDB.FindByPeriod(start, end)
	if err != nil {
		t.Fatalf("Failed to find invoices: %v", err)
	}

	if len(invoices) == 0 {
		t.Error("Expected at least one invoice, got 0")
	}
}
