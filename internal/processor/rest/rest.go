package rest

import (
	"app/internal/processor/rest/utils"
	"sync"
)

type REST struct {
	handler       *utils.Handler
	handlerRunner sync.Once

	controllers       *utils.Controllers
	controllersRunner sync.Once

	middlewares       *utils.Middlewares
	middlewaresRunner sync.Once
}

func Init() *REST {
	return &REST{}
}

func (r *REST) Handler() *utils.Handler {
	r.handlerRunner.Do(func() {
		r.handler = utils.InitHandler()
		r.handler.AddRoutesV1()
	})
	return r.handler
}

func (r *REST) Controllers() *utils.Controllers {
	r.controllersRunner.Do(func() {
		r.controllers = utils.InitControllers()
	})
	return r.controllers
}

func (r *REST) Middlewares() *utils.Middlewares {
	r.middlewaresRunner.Do(func() {
		r.middlewares = utils.InitMiddlewares()
	})
	return r.middlewares
}
