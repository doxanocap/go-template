package manager

import (
	"app/internal/manager/interfaces"
	IService "app/internal/manager/interfaces/service"
	"app/internal/service"
	"sync"
)

type ServiceManager struct {
	repo interfaces.IRepository
	proc interfaces.IProcessor

	auth       IService.IAuthService
	authRunner sync.Once

	storage       IService.IStorageService
	storageRunner sync.Once
}

func InitServiceManager(repo interfaces.IRepository, proc interfaces.IProcessor) *ServiceManager {
	return &ServiceManager{
		repo: repo,
		proc: proc,
	}
}

func (s *ServiceManager) Auth() IService.IAuthService {
	s.authRunner.Do(func() {
		s.auth = service.InitAuthService(s.repo)
	})
	return s.auth
}

func (s *ServiceManager) Storage() IService.IStorageService {
	s.storageRunner.Do(func() {
		s.storage = service.InitStorageService(s.repo, s.proc)
	})
	return s.storage
}
