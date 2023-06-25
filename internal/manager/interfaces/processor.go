package interfaces

import "app/internal/manager/interfaces/processor"

type IProcessor interface {
	REST() processor.IRESTProcessor
	Storage() processor.IStorageProcessor
}
