package processor

import (
	"app/internal/manager/interfaces/processor/rest"
)

type IRESTProcessor interface {
	Handler() rest.IHandlerManager
	Middlewares() rest.IMiddlewareManager
}
