package manager

import (
	IRepository "app/internal/manager/interfaces/repository"
	"app/internal/repository"
	"app/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
)

type RepositoryManager struct {
	pool *pgxpool.Pool

	user       IRepository.IUserRepository
	userRunner sync.Once

	storage       IRepository.IStorageRepository
	storageRunner sync.Once
}

func InitRepositoryManager(pool *pgxpool.Pool) *RepositoryManager {
	return &RepositoryManager{
		pool: pool,
	}
}

func (r *RepositoryManager) User() IRepository.IUserRepository {
	r.userRunner.Do(func() {
		r.user = repository.InitUserRepository(r.pool, logger.Log.Named("[REPOSITORY][USER]"))
	})
	return r.user
}

func (r *RepositoryManager) Storage() IRepository.IStorageRepository {
	r.storageRunner.Do(func() {
		r.storage = repository.InitStorageRepository(r.pool, logger.Log.Named("[REPOSITORY][STORAGE]"))
	})
	return r.storage
}
