package service

import (
	"app/internal/manager/interfaces"
	"app/internal/manager/interfaces/repository"
)

type AuthService struct {
	authRepo repository.IAutRepository
}

func InitAuthService(repo interfaces.IRepository) *AuthService {
	return &AuthService{
		authRepo: repo.Auth(),
	}
}
