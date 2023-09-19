package ws

import (
	"app/internal/manager/interfaces"
	IWS "app/internal/manager/interfaces/processor/ws"
	"go.uber.org/zap"
	"sync"
)

type WS struct {
	manager      interfaces.IManager
	poolProvider *PoolService

	clientService       IWS.IClientService
	clientServiceRunner sync.Once

	poolService       IWS.IPoolService
	poolServiceRunner sync.Once

	log *zap.Logger
}

func Init(manager interfaces.IManager, log *zap.Logger) *WS {
	return &WS{
		manager: manager,
		log:     log,
	}
}

func (ws *WS) Pool() IWS.IPoolService {
	ws.poolServiceRunner.Do(func() {
		ws.poolProvider = initPoolService(ws.manager, ws.log.Named("[POOL]"))
		ws.poolService = ws.poolProvider
	})
	return ws.poolService
}

func (ws *WS) Client() IWS.IClientService {
	ws.clientServiceRunner.Do(func() {
		ws.clientService = initClientService(ws.manager, ws.poolProvider, ws.log.Named("[CLIENT]"))
	})
	return ws.clientService
}
