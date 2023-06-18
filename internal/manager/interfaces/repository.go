package interfaces

import "app/internal/manager/interfaces/repository"

type IRepository interface {
	Auth() repository.IAuthRepository
	Storage() repository.IStorageRepository
}
