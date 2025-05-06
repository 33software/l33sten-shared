package rabbitmq

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher interface {
	Publish(ctx context.Context, routingKey string, payload any) error
	Close() error
}

type RabbitmqPublisher struct {
	conn     *amqp.Connection
	ch       *amqp.Channel
	exchange string
}

func NewPublisher(ctx context.Context, mqURL, exchange string, durable bool) (Publisher, error) {
	conn, err := amqp.DialConfig(mqURL, amqp.Config{Heartbeat: time.Second * 10})
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	if err := ch.ExchangeDeclare(
		exchange,
		"topic",
		durable,
		false,
		false,
		false,
		nil,
	); err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitmqPublisher{conn: conn, ch: ch, exchange: exchange}, nil
}

func (p *RabbitmqPublisher) Publish(ctx context.Context, routingKey string, payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return p.ch.PublishWithContext(
		ctx, p.exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
			Body:         body,
		})
}

func (p *RabbitmqPublisher) Close() error {
	p.ch.Close()
	return p.conn.Close()

}
