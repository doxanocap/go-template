package processor

import "app/internal/processor/rest"

type IRESTProcessor interface {
	Handler() *rest.Handler
	Controllers() *rest.Controllers
	Middlewares() *rest.Middlewares
}
