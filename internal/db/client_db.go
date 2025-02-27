package db

import (
	"database/sql"
	"super-shiharai-kun/internal/models"
)

type ClientDB struct {
	db *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {
	return &ClientDB{db: db}
}

func (r *ClientDB) Create(client *models.Client) error {
	query := `
        INSERT INTO clients (company_id, name, representative, phone, postal_code, address)
        VALUES (?, ?, ?, ?, ?, ?)
    `
	_, err := r.db.Exec(query,
		client.CompanyID,
		client.Name,
		client.Representative,
		client.Phone,
		client.PostalCode,
		client.Address,
	)
	return err
}

func (r *ClientDB) FindByCompanyID(companyID int) ([]models.Client, error) {
	query := `SELECT * FROM clients WHERE company_id = ?`
	rows, err := r.db.Query(query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		if err := rows.Scan(
			&client.ID,
			&client.CompanyID,
			&client.Name,
			&client.Representative,
			&client.Phone,
			&client.PostalCode,
			&client.Address,
		); err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}
