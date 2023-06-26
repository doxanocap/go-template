package interfaces

import "app/internal/manager/interfaces/repository"

type IRepository interface {
	User() repository.IUserRepository
	UserParams() repository.IUserParamsRepository
	Storage() repository.IStorageRepository
}
