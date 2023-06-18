package interfaces

type IManager interface {
	Processor() IProcessor
	Service() IService
	Repository() IRepository
}
