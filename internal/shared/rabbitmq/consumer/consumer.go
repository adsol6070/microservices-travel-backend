package consumer

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	QueueName         string
	AutoAck           bool
	RabbitMQURL       string
	ReconnectInterval time.Duration
}

type RabbitMQConsumer struct {
	Config  Config
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQConsumer(config Config) (*RabbitMQConsumer, error) {
	consumer := &RabbitMQConsumer{Config: config}
	err := consumer.connect()
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func (r *RabbitMQConsumer) connect() error {
	var err error
	r.Conn, err = amqp.Dial(r.Config.RabbitMQURL)
	if err != nil {
		return err
	}

	r.Channel, err = r.Conn.Channel()
	if err != nil {
		return err

	}

	return nil
}

func (r *RabbitMQConsumer) reconnect() error {
	log.Println("Reconnecting to RabbitMQ...")

	if r.Conn != nil {
		r.Conn.Close()
	}
	if r.Channel != nil {
		r.Channel.Close()
	}

	for {
		err := r.connect()
		if err == nil {
			log.Println("Reconnected successfully!")
			return nil
		}
		log.Println("Failed to reconnect:", err)
		time.Sleep(r.Config.ReconnectInterval)
	}
}

func (r *RabbitMQConsumer) Consume() (<-chan amqp.Delivery, error) {
	msgs, err := r.Channel.Consume(
		r.Config.QueueName, // queue
		"",                 // consumer tag
		r.Config.AutoAck,   // auto-acknowledgment
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // arguments
	)

	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (r *RabbitMQConsumer) AcknowledgeMessage(msg amqp.Delivery) error {
	if !r.Config.AutoAck {
		err := msg.Ack(false)
		if err != nil {
			log.Println("Failed to acknowledge message:", err)
			return err
		}
		log.Println("Message acknowledged.")
	}
	return nil
}

func (r *RabbitMQConsumer) GracefulShutdown() {
	// Listen for interrupt signals to shut down gracefully
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-sig
	log.Println("Graceful shutdown initiated...")

	// Close RabbitMQ connection and channel
	if r.Conn != nil {
		r.Conn.Close()
	}
	if r.Channel != nil {
		r.Channel.Close()
	}
	log.Println("RabbitMQ connection and channel closed.")
}
