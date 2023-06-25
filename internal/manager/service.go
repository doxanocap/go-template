package manager

import (
	"app/internal/manager/interfaces"
	IService "app/internal/manager/interfaces/service"
	"app/internal/service"
	"sync"
)

type ServiceManager struct {
	manager interfaces.IManager

	auth       IService.IAuthService
	authRunner sync.Once

	user       IService.IUserService
	userRunner sync.Once

	storage       IService.IStorageService
	storageRunner sync.Once
}

func InitServiceManager(manager interfaces.IManager) *ServiceManager {
	return &ServiceManager{
		manager: manager,
	}
}

func (s *ServiceManager) Auth() IService.IAuthService {
	s.authRunner.Do(func() {
		s.auth = service.InitAuthService(s.manager)
	})
	return s.auth
}

func (s *ServiceManager) User() IService.IUserService {
	s.userRunner.Do(func() {
		s.user = service.InitUserService(s.manager)
	})
	return s.user
}

func (s *ServiceManager) Storage() IService.IStorageService {
	s.storageRunner.Do(func() {
		s.storage = service.InitStorageService(s.manager)
	})
	return s.storage
}
