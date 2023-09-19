package interfaces

import "app/internal/manager/interfaces/processor"

type IProcessor interface {
	REST() processor.IRESTProcessor
	MsgBroker() processor.IMsgBrokerProcessor
	Mailer() processor.IMailerProcessor
	Storage() processor.IStorageProcessor
	WS() processor.IWSProcessor
}
