package runner

import (
	"app/internal/manager/interfaces"
	"app/internal/manager/interfaces/processor/rest"
	IMiddlewares "app/internal/manager/interfaces/processor/rest/middlewares"
	"app/internal/processor/rest/middlewares"
)

type Middlewares struct {
	authMiddleware IMiddlewares.IAuthMiddlewares
}

func InitMiddlewares(manager interfaces.IManager) rest.IMiddlewaresManager {
	return &Middlewares{
		authMiddleware: middlewares.InitAuthMiddleware(manager),
	}
}

func (m *Middlewares) Auth() IMiddlewares.IAuthMiddlewares {
	return m.authMiddleware
}
