package rest

import (
	"app/internal/manager/interfaces"
	"app/internal/manager/interfaces/processor/rest"
	"app/internal/processor/rest/handler"
	"app/internal/processor/rest/runner"
	"sync"
)

type REST struct {
	manager interfaces.IManager

	handler       rest.IHandlerManager
	handlerRunner sync.Once

	controllers       rest.IControllersManager
	controllersRunner sync.Once

	middlewares       rest.IMiddlewaresManager
	middlewaresRunner sync.Once
}

func Init(manager interfaces.IManager) *REST {
	return &REST{
		manager: manager,
	}
}

func (r *REST) Handler() rest.IHandlerManager {
	r.handlerRunner.Do(func() {
		r.handler = handler.InitHandler(r.Controllers())
	})
	return r.handler
}

func (r *REST) Controllers() rest.IControllersManager {
	r.controllersRunner.Do(func() {
		r.controllers = runner.InitControllers(r.manager)
	})
	return r.controllers
}

func (r *REST) Middlewares() rest.IMiddlewaresManager {
	r.middlewaresRunner.Do(func() {
		r.middlewares = runner.InitMiddlewares(r.manager)
	})
	return r.middlewares
}
