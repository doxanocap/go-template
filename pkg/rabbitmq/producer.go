package rabbitmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *MQClient) Send(ctx context.Context, qname, msg string) error {
	err := c.Ch.PublishWithContext(ctx,
		"",    // exchange
		qname, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	return err
}
