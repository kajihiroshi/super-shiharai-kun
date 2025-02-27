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
