package manager

import (
	IProcessor "app/internal/manager/interfaces/processor"
	"app/internal/processor/rest"
	"sync"
)

type ProcessorManager struct {
	rest       IProcessor.IRESTProcessor
	restRunner sync.Once
}

func InitProcessor() *ProcessorManager {
	return &ProcessorManager{}
}

func (p *ProcessorManager) REST() IProcessor.IRESTProcessor {
	p.restRunner.Do(func() {
		p.rest = rest.Init()
	})
	return p.rest
}
