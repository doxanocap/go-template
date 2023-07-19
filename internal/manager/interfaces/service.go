package interfaces

import "app/internal/manager/interfaces/service"

type IService interface {
	Storage() service.IStorageService
}
