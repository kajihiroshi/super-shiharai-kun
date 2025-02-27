package service

import (
	"super-shiharai-kun/internal/db"
	"super-shiharai-kun/internal/models"
)

type UserService struct {
	db *db.UserDB
}

func NewUserService(db *db.UserDB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.db.Create(user)
}
