package rabbitmq

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer interface {
	Start(ctx context.Context) error
	Close() error
}
type HandlerFunc func(ctx context.Context, body []byte) error

type RabbitmqConsumer struct {
	conn *amqp.Connection
	ch *amqp.Channel
	queue string
	handler HandlerFunc	
}

func NewConsumer(mqURL, exchange, qName string, handler HandlerFunc, routingKey ...string) (Consumer, error) {
	conn, err := amqp.Dial(mqURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	if err := ch.ExchangeDeclare(exchange, "topic", true, false, false, false, nil); err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	q, err := ch.QueueDeclare(qName, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	for _, key := range routingKey {
		if err := ch.QueueBind(q.Name, key, exchange, false, nil); err != nil {
			return nil, err
		}	
	}

	return &RabbitmqConsumer{conn: conn, ch: ch, queue: q.Name, handler: handler}, nil
}

func (c *RabbitmqConsumer) Start (ctx context.Context) error {
	//qos

	msgs, err := c.ch.Consume(c.queue, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go func () {
		for {
			select {
			case <- ctx.Done():
				return
			case msg, ok := <- msgs:
				if !ok {
					return
				}
				tCtx, cancel := context.WithTimeout(ctx, time.Second * 15)
				
				if err := c.handler(tCtx, msg.Body); err != nil {
					msg.Nack(false, true)
				} else {
					msg.Ack(false)
				}

				cancel()
			}
		}
	}()

	return nil
}

func (c *RabbitmqConsumer) Close() error{
	c.ch.Close()
	return c.conn.Close()
}