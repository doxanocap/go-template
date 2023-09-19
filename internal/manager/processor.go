package manager

import (
	"app/internal/manager/interfaces"
	IProcessor "app/internal/manager/interfaces/processor"
	"app/internal/processor/mailer"
	"app/internal/processor/msgbroker"
	"app/internal/processor/rest"
	"app/internal/processor/storage"
	"app/internal/processor/ws"
	"app/pkg/logger"
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

	wsProcessor       IProcessor.IWSProcessor
	wsProcessorRunner sync.Once
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
		p.msgBrokerProcessor = msgbroker.Init(p.msgBrokerProvider, logger.Log.Named("[PROCESSOR][MSG_BROKER]"))
	})
	return p.msgBrokerProcessor
}

func (p *ProcessorManager) Mailer() IProcessor.IMailerProcessor {
	p.mailerProcessorRunner.Do(func() {
		p.mailerProcessor = mailer.Init(p.mailerProvider, logger.Log.Named("[PROCESSOR][MAILER]"))
	})
	return p.mailerProcessor
}

func (p *ProcessorManager) Storage() IProcessor.IStorageProcessor {
	p.storageProcessorRunner.Do(func() {
		p.storageProcessor = storage.Init(p.storageProvider, logger.Log.Named("[PROCESSOR][STORAGE]"))
	})
	return p.storageProcessor
}

func (p *ProcessorManager) WS() IProcessor.IWSProcessor {
	p.wsProcessorRunner.Do(func() {
		p.wsProcessor = ws.Init(p.manager, logger.Log.Named("[PROCESSOR][WS]"))
	})
	return p.wsProcessor
}
