package rest

import (
	"sync"
)

type REST struct {
	handler       *Handler
	handlerRunner sync.Once

	controllers       *Controllers
	controllersRunner sync.Once

	middlewares       *Middlewares
	middlewaresRunner sync.Once
}

func Init() *REST {
	return &REST{}
}

func (r *REST) Handler() *Handler {
	r.handlerRunner.Do(func() {
		r.handler = InitHandler()
	})
	return r.handler
}

func (r *REST) Controllers() *Controllers {
	r.controllersRunner.Do(func() {
		r.controllers = InitControllers()
	})
	return r.controllers
}

func (r *REST) Middlewares() *Middlewares {
	r.middlewaresRunner.Do(func() {
		r.middlewares = InitMiddlewares()
	})
	return r.middlewares
}
