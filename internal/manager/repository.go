package manager

import (
	IRepository "app/internal/manager/interfaces/repository"
	"app/internal/repository"
	"app/pkg/logger"
	"gorm.io/gorm"
	"sync"
)

type RepositoryManager struct {
	conn *gorm.DB
	
	storage       IRepository.IStorageRepository
	storageRunner sync.Once
}

func InitRepositoryManager(conn *gorm.DB) *RepositoryManager {
	return &RepositoryManager{
		conn: conn,
	}
}

func (r *RepositoryManager) Storage() IRepository.IStorageRepository {
	r.storageRunner.Do(func() {
		r.storage = repository.InitStorageRepository(r.conn, logger.Log.Named("[REPOSITORY][STORAGE]"))
	})
	return r.storage
}
