package processor

import (
	IWS "app/internal/manager/interfaces/processor/ws"
)

type IWSProcessor interface {
	Pool() IWS.IPoolService
	Client() IWS.IClientService
}
