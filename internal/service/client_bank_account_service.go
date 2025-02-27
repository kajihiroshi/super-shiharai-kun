package service

import (
	"super-shiharai-kun/internal/db"
	"super-shiharai-kun/internal/models"
)

type ClientBankAccountService struct {
	db *db.ClientBankAccountDB
}

func NewClientBankAccountService(db *db.ClientBankAccountDB) *ClientBankAccountService {
	return &ClientBankAccountService{db: db}
}

func (s *ClientBankAccountService) CreateAccount(account *models.ClientBankAccount) error {
	return s.db.Create(account)
}

func (s *ClientBankAccountService) GetAccountsByClientID(clientID int) ([]models.ClientBankAccount, error) {
	return s.db.FindByClientID(clientID)
}
