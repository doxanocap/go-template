package manager

import (
	"app/internal/manager/interfaces"
	"app/internal/manager/interfaces/processor"
	"app/pkg/redis"
	"app/pkg/smtp"
	_ "github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/gorm"
	"sync"
)

type Manager struct {
	conn              *gorm.DB
	cacheConn         *redis.Conn
	storageProvider   processor.IStorageProvider
	msgBrokerProvider processor.IMsgBrokerProvider
	mailerProvider    processor.IMailerProvider

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
		m.repository = InitRepositoryManager(m.conn)
	})
	return m.repository
}

func (m *Manager) Service() interfaces.IService {
	m.serviceRunner.Do(func() {
		m.service = InitServiceManager(m)
	})
	return m.service
}

func (m *Manager) Processor() interfaces.IProcessor {
	m.processorRunner.Do(func() {
		m.processor = InitProcessor(m, m.storageProvider)
	})
	return m.processor
}

func (m *Manager) SetConn(conn *gorm.DB) {
	m.conn = conn
}

func (m *Manager) SetStorageProvider(storageProvider processor.IStorageProvider) {
	m.storageProvider = storageProvider
}

func (m *Manager) SetMsgBroker(msgBrokerProvider processor.IMsgBrokerProvider) {
	m.msgBrokerProvider = msgBrokerProvider
}

func (m *Manager) SetCacheConnection(cacheConn *redis.Conn) {
	m.cacheConn = cacheConn
}
func (m *Manager) SetMailer(mailer *smtp.SMTP) {
	m.mailerProvider = mailer
}
