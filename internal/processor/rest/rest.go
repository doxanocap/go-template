package rest

import (
	"app/internal/manager/interfaces"
	"app/internal/manager/interfaces/processor/rest"
	"app/internal/processor/rest/handler"
	"app/internal/processor/rest/middlewares"
	"app/pkg/logger"
	"sync"
)

type REST struct {
	manager interfaces.IManager

	handler       rest.IHandlerManager
	handlerRunner sync.Once

	middlewares       rest.IMiddlewareManager
	middlewaresRunner sync.Once
}

func Init(manager interfaces.IManager) *REST {
	return &REST{
		manager: manager,
	}
}

func (r *REST) Handler() rest.IHandlerManager {
	r.handlerRunner.Do(func() {
		r.handler = handler.InitHandler(r.manager, logger.Log.Named("[HANDLER]"))
	})
	return r.handler
}

func (r *REST) Middlewares() rest.IMiddlewareManager {
	r.middlewaresRunner.Do(func() {
		r.middlewares = middlewares.InitMiddlewares(r.manager)
	})
	return r.middlewares
}
