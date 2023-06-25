package interfaces

import "app/internal/manager/interfaces/service"

type IService interface {
	Auth() service.IAuthService
	User() service.IUserService
	Storage() service.IStorageService
}
