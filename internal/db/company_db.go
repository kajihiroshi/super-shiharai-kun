package db

import (
	"database/sql"
	"super-shiharai-kun/internal/models"
)

type CompanyDB struct {
	db *sql.DB
}

func NewCompanyDB(db *sql.DB) *CompanyDB {
	return &CompanyDB{db: db}
}

func (r *CompanyDB) Create(company *models.Company) error {
	query := `INSERT INTO companies (name, representative, phone, postal_code, address) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, company.Name, company.Representative, company.Phone, company.PostalCode, company.Address)
	return err
}

func (r *CompanyDB) FindByID(id int) (*models.Company, error) {
	query := `SELECT * FROM companies WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var company models.Company
	if err := row.Scan(
		&company.ID,
		&company.Name,
		&company.Representative,
		&company.Phone,
		&company.PostalCode,
		&company.Address,
	); err != nil {
		return nil, err
	}
	return &company, nil
}
