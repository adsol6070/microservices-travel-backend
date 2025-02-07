package main

import (
	"log"
	"microservices-travel-backend/internal/email-service/infrastructure/sendgrid"
	"microservices-travel-backend/internal/email-service/services"
	"microservices-travel-backend/internal/shared/rabbitmq/consumer"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	consumerConfig := consumer.Config{
		QueueName:         "emailQueue",
		AutoAck:           false,
		RabbitMQURL:       os.Getenv("RABBITMQ_URL"),
		ReconnectInterval: 6 * time.Second,
	}

	rabbitMQConsumer, err := consumer.NewRabbitMQConsumer(consumerConfig)
	if err != nil {
		log.Fatal("Failed to initialize RabbitMQ consumer:", err)
	}

	sendGridClient := sendgrid.NewSendGridClient()
	emailService := services.NewEmailService(sendGridClient)

	msgs, err := rabbitMQConsumer.Consume()
	if err != nil {
		log.Fatal("Failed to consume messages:", err)
	}

	log.Println("Email Service is running and waiting for messages...")

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		log.Println("Received shutdown signal. Cleaning up...")
		rabbitMQConsumer.GracefulShutdown()
		os.Exit(0)
	}()

	for msg := range msgs {
		emailService.ProcessEmailMessage(msg)
		rabbitMQConsumer.AcknowledgeMessage(msg)
	}
}
