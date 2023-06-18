package middlewares

import "app/internal/manager/interfaces/processor/rest"

type Middlewares struct {
}

func InitMiddlewares() rest.IMiddlewaresManager {
	return &Middlewares{}
}
