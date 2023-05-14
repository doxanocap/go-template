package manager

import (
	IRepo "app/internal/manager/interfaces/repository"
	"app/internal/repository"
	"app/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
)

type RepositoryManager struct {
	pool *pgxpool.Pool

	auth     IRepo.IAutRepository
	authInit sync.Once
}

func InitRepositoryManager(pool *pgxpool.Pool) *RepositoryManager {
	return &RepositoryManager{
		pool: pool,
	}
}

func (r *RepositoryManager) Auth() IRepo.IAutRepository {
	r.authInit.Do(func() {
		r.auth = repository.InitAuthRepository(r.pool, logger.Log.Named("[REPOSITORY][AUTH]"))
	})
	return r.auth
}
