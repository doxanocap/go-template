package manager

import (
	"app/internal/manager/interfaces"
	IProcessor "app/internal/manager/interfaces/processor"
	"app/internal/processor/rest"
	"app/internal/processor/storage"
	"sync"
)

type ProcessorManager struct {
	manager         interfaces.IManager
	storageProvider IProcessor.IStorageProvider

	rest       IProcessor.IRESTProcessor
	restRunner sync.Once

	storageProcessor       IProcessor.IStorageProcessor
	storageProcessorRunner sync.Once
}

func InitProcessor(manager interfaces.IManager, storageProvider IProcessor.IStorageProvider) *ProcessorManager {
	return &ProcessorManager{
		manager:         manager,
		storageProvider: storageProvider,
	}
}

func (p *ProcessorManager) REST() IProcessor.IRESTProcessor {
	p.restRunner.Do(func() {
		p.rest = rest.Init(p.manager)
	})
	return p.rest
}

func (p *ProcessorManager) Storage() IProcessor.IStorageProcessor {
	p.storageProcessorRunner.Do(func() {
		p.storageProcessor = storage.Init(p.storageProvider)
	})
	return p.storageProcessor
}
