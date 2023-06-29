package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type ConsumerParams struct {
	QueueName string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

func (c *MQClient) Consume(obj ConsumerParams) error {
	messages, err := c.Ch.Consume(obj.QueueName, "", obj.AutoAck, obj.Exclusive, obj.NoLocal, obj.NoWait, obj.Args)
	if err != nil {
		return err
	}

	for message := range messages {
		log.Printf("[CONSUMER] received a message: %s", message.Body)
	}
	return nil
}
