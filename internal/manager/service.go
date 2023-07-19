package manager

import (
	"app/internal/manager/interfaces"
	IService "app/internal/manager/interfaces/service"
	"app/internal/service"
	"sync"
)

type ServiceManager struct {
	manager interfaces.IManager

	storage       IService.IStorageService
	storageRunner sync.Once
}

func InitServiceManager(manager interfaces.IManager) *ServiceManager {
	return &ServiceManager{
		manager: manager,
	}
}

func (s *ServiceManager) Storage() IService.IStorageService {
	s.storageRunner.Do(func() {
		s.storage = service.InitStorageService(s.manager)
	})
	return s.storage
}
