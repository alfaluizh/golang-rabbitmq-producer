package rabbitmq

import (
	"context"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	conn       *amqp.Connection
	queue      string
	replyQueue string
}

func NewProducer(url, queue, replyQueue string) *Producer {
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	_, err = ch.QueueDeclare(replyQueue, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	return &Producer{
		conn:       conn,
		queue:      queue,
		replyQueue: replyQueue,
	}
}

func (p *Producer) Send(message string) (string, error) {
	ch, err := p.conn.Channel()
	if err != nil {
		return "", err
	}
	defer ch.Close()

	correlationID := uuid.NewString()

	err = ch.PublishWithContext(context.Background(), "", p.queue, false, false, amqp.Publishing{
		ContentType:   "text/plain",
		CorrelationId: correlationID,
		ReplyTo:       p.replyQueue,
		Body:          []byte(message),
	})
	if err != nil {
		return "", err
	}

	msgs, err := ch.Consume(p.replyQueue, "", false, false, false, false, nil)
	if err != nil {
		return "", err
	}

	timeout := time.After(10 * time.Second)

	for {
		select {
		case msg := <-msgs:
			if msg.CorrelationId == correlationID {
				msg.Ack(false)
				return string(msg.Body), nil
			}
			msg.Nack(false, true)
		case <-timeout:
			return "", context.DeadlineExceeded
		}
	}
}
