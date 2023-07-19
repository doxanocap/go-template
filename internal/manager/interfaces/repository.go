package interfaces

import "app/internal/manager/interfaces/repository"

type IRepository interface {
	Storage() repository.IStorageRepository
}
