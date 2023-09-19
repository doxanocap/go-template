package interfaces

import "app/internal/model"

type IManager interface {
	Processor() IProcessor
	Service() IService
	Repository() IRepository
	
	Cfg() *model.Config
}
