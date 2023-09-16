package rabbitmq

import (
	"app/internal/config"
	"github.com/doxanocap/pkg/lg"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MQClient struct {
	Ch     *amqp.Channel
	Queues map[string]amqp.Queue
}

func InitConnection(cfg *config.Cfg) *MQClient {
	conn, err := amqp.Dial(cfg.RabbitDSN)
	if err != nil {
		return nil
	}

	defer func() {
		err = conn.Close()
	}()

	ch, err := conn.Channel()
	if err != nil {
		lg.Fatalf("connection to MQClient: %v", err)
	}

	defer func() {
		err = ch.Close()
	}()

	return &MQClient{
		Ch:     ch,
		Queues: map[string]amqp.Queue{},
	}
}

type QueueParams struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

func (c *MQClient) NewQueue(obj QueueParams) error {
	q, err := c.Ch.QueueDeclare(obj.Name, obj.Durable, obj.AutoDelete, obj.Exclusive, obj.NoWait, obj.Args)
	if err != nil {
		return err
	}
	c.Queues[obj.Name] = q
	return nil
}
