package rest

import (
	"app/internal/manager/interfaces"
	"app/internal/manager/interfaces/processor/rest"
	"app/internal/processor/rest/controllers"
	"app/internal/processor/rest/handler"
	"app/internal/processor/rest/middlewares"
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
		r.handler = handler.InitHandler(r.manager)
	})
	return r.handler
}

func (r *REST) Controllers() rest.IControllersManager {
	r.controllersRunner.Do(func() {
		r.controllers = controllers.InitControllers(r.manager)
	})
	return r.controllers
}

func (r *REST) Middlewares() rest.IMiddlewaresManager {
	r.middlewaresRunner.Do(func() {
		r.middlewares = middlewares.InitMiddlewares()
	})
	return r.middlewares
}
