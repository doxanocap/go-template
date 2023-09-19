package manager

import (
	"app/internal/manager/interfaces"
	"app/internal/repository"
	"app/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
)

type RepositoryManager struct {
	pool *pgxpool.Pool

	user       interfaces.IUserRepository
	userRunner sync.Once

	userParams       interfaces.IUserParamsRepository
	userParamsRunner sync.Once

	storage       interfaces.IStorageRepository
	storageRunner sync.Once
}

func InitRepositoryManager(pool *pgxpool.Pool) *RepositoryManager {
	return &RepositoryManager{
		pool: pool,
	}
}

func (r *RepositoryManager) User() interfaces.IUserRepository {
	r.userRunner.Do(func() {
		r.user = repository.InitUserRepository(r.pool, logger.Log.Named("[REPOSITORY][USER]"))
	})
	return r.user
}

func (r *RepositoryManager) UserParams() interfaces.IUserParamsRepository {
	r.userParamsRunner.Do(func() {
		r.userParams = repository.InitUserParamsRepository(r.pool, logger.Log.Named("[REPOSITORY][USER_PARAMS]"))
	})
	return r.userParams
}

func (r *RepositoryManager) Storage() interfaces.IStorageRepository {
	r.storageRunner.Do(func() {
		r.storage = repository.InitStorageRepository(r.pool, logger.Log.Named("[REPOSITORY][STORAGE]"))
	})
	return r.storage
}
