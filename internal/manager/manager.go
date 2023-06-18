package manager

import (
	"app/internal/manager/interfaces"
	"app/internal/manager/interfaces/processor"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/pgxpool"
	"sync"
)

type Manager struct {
	pool            *pgxpool.Pool
	storageProvider processor.IStorageProvider

	service       interfaces.IService
	serviceRunner sync.Once

	repository       interfaces.IRepository
	repositoryRunner sync.Once

	processor       interfaces.IProcessor
	processorRunner sync.Once
}

func InitManager() *Manager {
	return &Manager{}
}

func (m *Manager) Repository() interfaces.IRepository {
	m.repositoryRunner.Do(func() {
		m.repository = InitRepositoryManager(m.pool)
	})
	return m.repository
}

func (m *Manager) Service() interfaces.IService {
	m.serviceRunner.Do(func() {
		m.service = InitServiceManager(m.Repository(), m.Processor())
	})
	return m.service
}

func (m *Manager) Processor() interfaces.IProcessor {
	m.processorRunner.Do(func() {
		m.processor = InitProcessor(m, m.storageProvider)
	})
	return m.processor
}

func (m *Manager) SetPool(pool *pgxpool.Pool) {
	m.pool = pool
}

func (m *Manager) SetStorageProvider(storageProvider processor.IStorageProvider) {
	m.storageProvider = storageProvider
}
