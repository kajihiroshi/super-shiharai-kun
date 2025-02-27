package service

import (
	"super-shiharai-kun/internal/db"
	"super-shiharai-kun/internal/models"
)

type ClientService struct {
	db *db.ClientDB
}

func NewClientService(db *db.ClientDB) *ClientService {
	return &ClientService{db: db}
}

func (s *ClientService) CreateClient(client *models.Client) error {
	return s.db.Create(client)
}

func (s *ClientService) GetClientsByCompanyID(companyID int) ([]models.Client, error) {
	return s.db.FindByCompanyID(companyID)
}
