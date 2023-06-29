package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

type MQClient struct {
	Ch     *amqp.Channel
	Queues map[string]amqp.Queue
}

func Connect() (client *MQClient, err error) {
	conn, err := amqp.Dial(viper.GetString("RABBIT_MQ_DSN"))
	if err != nil {
		return
	}

	defer func() {
		err = conn.Close()
	}()

	ch, err := conn.Channel()
	if err != nil {
		return
	}

	defer func() {
		err = ch.Close()
	}()

	client = &MQClient{
		Ch:     ch,
		Queues: map[string]amqp.Queue{},
	}
	return
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
