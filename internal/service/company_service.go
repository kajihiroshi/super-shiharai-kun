package service

import (
	"super-shiharai-kun/internal/db"
	"super-shiharai-kun/internal/models"
)

type CompanyService struct {
	db *db.CompanyDB
}

func NewCompanyService(db *db.CompanyDB) *CompanyService {
	return &CompanyService{db: db}
}

func (s *CompanyService) CreateCompany(company *models.Company) error {
	return s.db.Create(company)
}
