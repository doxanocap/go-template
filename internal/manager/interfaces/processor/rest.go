package processor

import (
	"app/internal/processor/rest/utils"
)

type IRESTProcessor interface {
	Handler() *utils.Handler
	Controllers() *utils.Controllers
	Middlewares() *utils.Middlewares
}
