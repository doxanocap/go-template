package service

import (
	"app/internal/manager/interfaces"
	"app/internal/model"
)

type UserService struct {
	repository interfaces.IRepository
}

func InitUserService(repository interfaces.IRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (as *UserService) Create(body model.SignUp) model.AuthResponse {
	
}
