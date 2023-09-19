package processor

import (
	"app/pkg/rabbitmq"
	"context"
)

type IMsgBrokerProcessor interface {
	Send(ctx context.Context, qname, msg string) error
	Consume(ctx context.Context, obj rabbitmq.ConsumerParams) error
	NewQueue(obj rabbitmq.QueueParams) error
}

type IMsgBrokerProvider interface {
	Send(ctx context.Context, qname, msg string) error
	Consume(obj rabbitmq.ConsumerParams) error
	NewQueue(obj rabbitmq.QueueParams) error
}
