package external

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Publisher struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func (p *Publisher) Close() {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.connection != nil {
		p.connection.Close()
	}
}

func (p *Publisher) Publish(exchange string, routingKey string, body []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := p.channel.PublishWithContext(
		ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	return err
}

func NewPublisher(amqpURI string) Publisher {
	connection, err := amqp.DialConfig(amqpURI, amqp.Config{Properties: amqp.NewConnectionProperties()})
	if err != nil {
		log.Fatalf("Could not connect to rabbitmq: %v", err)
	}
	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("Could not open a channel: %v", err)
	}
	return Publisher{
		connection: connection,
		channel:    channel,
	}
}
