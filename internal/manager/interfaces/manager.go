package interfaces

type Manager interface {
	Processor() IProcessor
	Service() IService
	Repo() IRepository
}
