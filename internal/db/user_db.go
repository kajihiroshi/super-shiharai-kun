package db

import (
	"database/sql"
	"super-shiharai-kun/internal/models"
)

type UserDB struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{db: db}
}

func (r *UserDB) Create(user *models.User) error {
	query := `INSERT INTO users (company_id, name, email, password) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, user.CompanyID, user.Name, user.Email, user.Password)
	return err
}

func (r *UserDB) FindByEmail(email string) (*models.User, error) {
	query := `SELECT * FROM users WHERE email = ?`
	row := r.db.QueryRow(query, email)

	var user models.User
	if err := row.Scan(
		&user.ID,
		&user.CompanyID,
		&user.Name,
		&user.Email,
		&user.Password,
	); err != nil {
		return nil, err
	}
	return &user, nil
}
