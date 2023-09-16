package manager

import (
	"app/internal/manager/interfaces"
	"app/internal/service"
	"sync"
)

type ServiceManager struct {
	manager interfaces.IManager

	auth       interfaces.IAuthService
	authRunner sync.Once

	user       interfaces.IUserService
	userRunner sync.Once

	storage       interfaces.IStorageService
	storageRunner sync.Once
}

func InitServiceManager(manager interfaces.IManager) *ServiceManager {
	return &ServiceManager{
		manager: manager,
	}
}

func (s *ServiceManager) Auth() interfaces.IAuthService {
	s.authRunner.Do(func() {
		s.auth = service.InitAuthService(s.manager)
	})
	return s.auth
}

func (s *ServiceManager) User() interfaces.IUserService {
	s.userRunner.Do(func() {
		s.user = service.InitUserService(s.manager)
	})
	return s.user
}

func (s *ServiceManager) Storage() interfaces.IStorageService {
	s.storageRunner.Do(func() {
		s.storage = service.InitStorageService(s.manager)
	})
	return s.storage
}
