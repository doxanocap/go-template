package manager

import (
	"app/internal/manager/interfaces"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/pgxpool"
	"sync"
)

var repo *RepositoryManager

type Manager struct {
	pool *pgxpool.Pool

	service       interfaces.IService
	serviceRunner sync.Once

	repository       interfaces.IRepository
	repositoryRunner sync.Once

	processor       interfaces.IProcessor
	processorRunner sync.Once
}

func InitManager(pool *pgxpool.Pool) *Manager {
	return &Manager{
		pool: pool,
	}
}

func (m *Manager) Repository() interfaces.IRepository {
	m.repositoryRunner.Do(func() {
		repo = InitRepositoryManager(m.pool)
		m.repository = repo
	})
	return m.repository
}

func (m *Manager) Service() interfaces.IService {
	m.serviceRunner.Do(func() {
		m.service = InitServiceManager(repo)
	})
	return m.service
}

func (m *Manager) Processor() interfaces.IProcessor {
	m.processorRunner.Do(func() {
		m.processor = InitProcessor()
	})
	return m.processor
}
