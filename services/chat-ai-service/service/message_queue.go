package service

import (
	"fmt"
	"log"

	"github.com/crm/chat-ai-service/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageQueue struct {
	config     *config.Config
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewMessageQueue(cfg *config.Config) (*MessageQueue, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s%s",
		cfg.RabbitMQUser,
		cfg.RabbitMQPassword,
		cfg.RabbitMQHost,
		cfg.RabbitMQPort,
		cfg.RabbitMQVHost,
	)

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	// Declare queues
	_, err = ch.QueueDeclare(
		"chat_messages", // queue name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	return &MessageQueue{
		config:     cfg,
		connection: conn,
		channel:    ch,
	}, nil
}

func (mq *MessageQueue) StartConsuming() {
	msgs, err := mq.channel.Consume(
		"chat_messages", // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		log.Printf("Failed to register consumer: %v", err)
		return
	}

	log.Println("Starting to consume messages from RabbitMQ...")

	for msg := range msgs {
		log.Printf("Received message: %s", msg.Body)
		// Process message here
	}
}

func (mq *MessageQueue) Close() {
	if mq.channel != nil {
		mq.channel.Close()
	}
	if mq.connection != nil {
		mq.connection.Close()
	}
}
