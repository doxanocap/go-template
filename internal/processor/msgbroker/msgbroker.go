package msgbroker

import (
	"app/internal/manager/interfaces/processor"
	"app/pkg/logger"
	"app/pkg/rabbitmq"
	"context"
	"fmt"
)

type MsgBroker struct {
	provider processor.IMsgBrokerProvider
}

func (mb *MsgBroker) Send(ctx context.Context, qname, msg string) error {
	return mb.provider.Send(ctx, qname, msg)
}

func (mb *MsgBroker) Consume(ctx context.Context, obj rabbitmq.ConsumerParams) error {
	for {
		select {
		case <-ctx.Done():
			logger.Log.Info(fmt.Sprintf("consumer stopped listening to: %s", obj.QueueName))
			return nil
		default:
			err := mb.provider.Consume(obj)
			if err != nil {
				return err
			}
		}
	}
}

func (mb *MsgBroker) NewQueue(obj rabbitmq.QueueParams) error {
	return mb.provider.NewQueue(obj)
}

func Init(provider processor.IMsgBrokerProvider) *MsgBroker {
	return &MsgBroker{
		provider: provider,
	}
}
