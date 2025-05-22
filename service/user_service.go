package service

import (
	"wire-gorm-server/model"
	"wire-gorm-server/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Return all users
func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAll()
}
