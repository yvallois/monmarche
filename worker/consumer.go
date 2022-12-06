package worker

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type HandlerFunc func(string) bool

type Consumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	handler    HandlerFunc
	queueName  string
}

func (c *Consumer) Run() {
	msgs, err := c.channel.Consume(
		c.queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Could not consume: %v", err)
	}
	var forever chan struct{}

	go func() {
		for msg := range msgs {
			go c.processMessage(msg)
		}
	}()
	<-forever
}

func (c *Consumer) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.connection != nil {
		c.connection.Close()
	}
}

func (c *Consumer) processMessage(d amqp.Delivery) {
	var err error
	result := c.handler(string(d.Body))
	if result {
		err = d.Ack(false)
	} else {
		err = d.Nack(false, false)
	}
	if err != nil {
		log.Fatalf("Could ack the message: %v", err)
	}
}

func NewConsumer(amqpURI string, queueName string, handler HandlerFunc) Consumer {
	connection, err := amqp.DialConfig(amqpURI, amqp.Config{Properties: amqp.NewConnectionProperties()})
	if err != nil {
		log.Fatalf("Could not connect to rabbitmq: %v", err)
	}
	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("Could not open a channel: %v", err)
	}
	err = channel.Qos(10, 0, false)
	if err != nil {
		log.Fatalf("Could not set QoS: %v", err)
	}
	return Consumer{
		connection: connection,
		channel:    channel,
		handler:    handler,
		queueName:  queueName,
	}
}
