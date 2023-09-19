package middlewares

import (
	"app/internal/manager/interfaces"
	"app/internal/manager/interfaces/processor/rest"
)

type Middlewares struct {
	manager interfaces.IManager
}

func InitMiddlewares(manager interfaces.IManager) rest.IMiddlewareManager {
	return &Middlewares{
		manager: manager,
	}
}
