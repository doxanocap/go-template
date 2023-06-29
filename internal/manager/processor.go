package manager

import (
	"app/internal/manager/interfaces"
	IProcessor "app/internal/manager/interfaces/processor"
	"app/internal/processor/msgbroker"
	"app/internal/processor/rest"
	"app/internal/processor/storage"
	"sync"
)

type ProcessorManager struct {
	manager           interfaces.IManager
	storageProvider   IProcessor.IStorageProvider
	msgBrokerProvider IProcessor.IMsgBrokerProvider
	mailerProvider    IProcessor.IMailerProvider

	restProcessor       IProcessor.IRESTProcessor
	restProcessorRunner sync.Once

	storageProcessor       IProcessor.IStorageProcessor
	storageProcessorRunner sync.Once

	msgBrokerProcessor       IProcessor.IMsgBrokerProcessor
	msgBrokerProcessorRunner sync.Once

	mailerProcessor       IProcessor.IMailerProcessor
	mailerProcessorRunner sync.Once
}

func InitProcessor(manager interfaces.IManager, storageProvider IProcessor.IStorageProvider) *ProcessorManager {
	return &ProcessorManager{
		manager:         manager,
		storageProvider: storageProvider,
	}
}

func (p *ProcessorManager) REST() IProcessor.IRESTProcessor {
	p.restProcessorRunner.Do(func() {
		p.restProcessor = rest.Init(p.manager)
	})
	return p.restProcessor
}

func (p *ProcessorManager) MsgBroker() IProcessor.IMsgBrokerProcessor {
	p.msgBrokerProcessorRunner.Do(func() {
		p.msgBrokerProcessor = msgbroker.Init(p.msgBrokerProvider)
	})
	return p.msgBrokerProcessor
}

func (p *ProcessorManager) Mailer() IProcessor.IMailerProcessor {
	p.mailerProcessorRunner.Do(func() {
		p.mailerProcessor = p.mailerProvider
	})
	return p.mailerProcessor
}

func (p *ProcessorManager) Storage() IProcessor.IStorageProcessor {
	p.storageProcessorRunner.Do(func() {
		p.storageProcessor = storage.Init(p.storageProvider)
	})
	return p.storageProcessor
}
