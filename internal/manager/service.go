package manager

import (
	"app/internal/manager/interfaces/service"
	service2 "app/internal/service"
	"sync"
)

type ServiceManager struct {
	repo *RepositoryManager

	auth       service.IAuthService
	authRunner sync.Once
}

func InitServiceManager(repo *RepositoryManager) *ServiceManager {
	return &ServiceManager{
		repo: repo,
	}
}

func (s *ServiceManager) Auth() service.IAuthService {
	s.authRunner.Do(func() {
		s.auth = service2.InitAuthService(s.repo)
	})
	return s.auth
}
