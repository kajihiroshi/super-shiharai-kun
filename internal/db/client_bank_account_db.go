package db

import (
	"database/sql"
	"super-shiharai-kun/internal/models"
)

type ClientBankAccountDB struct {
	db *sql.DB
}

func NewClientBankAccountDB(db *sql.DB) *ClientBankAccountDB {
	return &ClientBankAccountDB{db: db}
}

func (r *ClientBankAccountDB) Create(account *models.ClientBankAccount) error {
	query := `
        INSERT INTO client_bank_accounts (client_id, bank_name, branch_name, account_number, account_name)
        VALUES (?, ?, ?, ?, ?)
    `
	_, err := r.db.Exec(query,
		account.ClientID,
		account.BankName,
		account.BranchName,
		account.AccountNumber,
		account.AccountName,
	)
	return err
}

func (r *ClientBankAccountDB) FindByClientID(clientID int) ([]models.ClientBankAccount, error) {
	query := `SELECT * FROM client_bank_accounts WHERE client_id = ?`
	rows, err := r.db.Query(query, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []models.ClientBankAccount
	for rows.Next() {
		var account models.ClientBankAccount
		if err := rows.Scan(
			&account.ID,
			&account.ClientID,
			&account.BankName,
			&account.BranchName,
			&account.AccountNumber,
			&account.AccountName,
		); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
