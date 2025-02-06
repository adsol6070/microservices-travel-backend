package rabbitmq

import (
	"context"
	"errors"
	"log"
	"os"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (client *Client) Publish(message []byte, queueName, addr string) error {
	queue := New(queueName, addr)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	errChan := make(chan error, 1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			case <-time.After(time.Second * 2):
				if err := queue.Push(message); err != nil {
					queue.errlog.Printf("Push failed: %v", err)
				} else {
					queue.infolog.Println("Push succeeded")
					errChan <- nil
					return
				}
			}
		}
	}()

	err := <-errChan
	if err != nil {
		queue.errlog.Printf("Failed to publish message: %v", err)
	} else {
		queue.infolog.Println("Message published successfully.")
	}

	if err := queue.Close(); err != nil {
		queue.errlog.Printf("Error closing the queue: %v", err)
	}

	return err
}

func (client *Client) consume(done chan struct{}) {
	client.m.Lock()
	if !client.isReady {
		client.m.Unlock()
		client.errlog.Println("Client not connected. Consume aborted.")
		close(done)
		return
	}
	client.m.Unlock()

	// Allow some time for the connection to establish before starting consumption
	<-time.After(time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*25)
	defer cancel()

	deliveries, err := client.Consume()
	if err != nil {
		client.errlog.Printf("could not start consuming: %s\n", err)
		return
	}

	// This channel will receive a notification when a channel closed event happens.
	// Make sure it's buffered to prevent deadlocks.
	chClosedCh := make(chan *amqp.Error, 1)
	client.channel.NotifyClose(chClosedCh)

loop:
	for {
		select {
		case <-ctx.Done():
			err := client.Close()
			if err != nil {
				client.errlog.Printf("close failed: %s\n", err)
			}
			break loop

		case amqErr := <-chClosedCh:
			// Log when the AMQP channel is closed due to abnormal shutdown
			client.errlog.Printf("AMQP Channel closed due to: %s\n", amqErr)

			// Attempt to re-consume after channel closure
			deliveries, err = client.Consume()
			if err != nil {
				// Re-attempt consumption until successful or timeout occurs
				client.errlog.Println("error trying to consume, will retry...")
				select {
				case <-time.After(time.Second * 5): // Retry interval
				case <-ctx.Done():
					break loop
				}
				continue
			}

			// Reset the channel to listen for future closures
			chClosedCh = make(chan *amqp.Error, 1)
			client.channel.NotifyClose(chClosedCh)

		case delivery := <-deliveries:
			// Log the received message and acknowledge it
			client.infolog.Printf("received message: %s\n", delivery.Body)

			if err := delivery.Ack(false); err != nil {
				client.errlog.Printf("error acknowledging message: %s\n", err)
				// You can add logic to requeue or log more details if needed.
			}

			// Controlled delay for message consumption
			<-time.After(time.Second * 2)
		}
	}

	// Final cleanup when the loop exits
	close(done)
}

type Client struct {
	m               *sync.Mutex
	queueName       string
	infolog         *log.Logger
	errlog          *log.Logger
	connection      *amqp.Connection
	channel         *amqp.Channel
	done            chan bool
	notifyConnClose chan *amqp.Error
	notifyChanClose chan *amqp.Error
	notifyConfirm   chan amqp.Confirmation
	isReady         bool
}

const (
	reconnectDelay = 5 * time.Second

	reInitDelay = 2 * time.Second

	resendDelay = 5 * time.Second
)

var (
	errNotConnected  = errors.New("not connected to a server")
	errAlreadyClosed = errors.New("already closed: not connected to the server")
	errShutdown      = errors.New("client is shutting down")
)

func New(queueName, addr string) *Client {
	client := Client{
		m:         &sync.Mutex{},
		infolog:   log.New(os.Stdout, "[INFO]", log.LstdFlags|log.Lmsgprefix),
		errlog:    log.New(os.Stderr, "[ERROR]", log.LstdFlags|log.Lmsgprefix),
		queueName: queueName,
		done:      make(chan bool),
	}

	go client.handleReconnect(addr)

	return &client
}

func (client *Client) handleReconnect(addr string) {
	for {
		client.m.Lock()
		client.isReady = false
		client.m.Unlock()

		client.infolog.Println("attempting to connect")

		conn, err := client.connect(addr)
		if err != nil {
			client.errlog.Println("failed to connect. Retrying...")

			select {
			case <-client.done:
				return
			case <-time.After(reconnectDelay):
			}
			continue
		}

		if done := client.handleReInit(conn); done {
			break
		}
	}
}

func (client *Client) connect(addr string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(addr)
	if err != nil {
		return nil, err
	}

	client.changeConnection(conn)
	client.infolog.Println("connected")
	return conn, nil
}

func (client *Client) init(conn *amqp.Connection) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	err = ch.Confirm(false)
	if err != nil {
		return err
	}
	_, err = ch.QueueDeclare(
		client.queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	client.changeChannel(ch)
	client.m.Lock()
	client.isReady = true
	client.m.Unlock()
	client.infolog.Println("client init done")

	return nil
}

func (client *Client) changeConnection(connection *amqp.Connection) {
	client.connection = connection
	client.notifyConnClose = make(chan *amqp.Error, 1)
	client.connection.NotifyClose(client.notifyConnClose)
}

func (client *Client) changeChannel(channel *amqp.Channel) {
	client.channel = channel
	client.notifyChanClose = make(chan *amqp.Error, 1)
	client.notifyConfirm = make(chan amqp.Confirmation, 1)
	client.channel.NotifyClose(client.notifyChanClose)
	client.channel.NotifyPublish(client.notifyConfirm)
}

func (client *Client) Push(data []byte) error {
	client.m.Lock()
	if !client.isReady {
		client.m.Unlock()
		return errNotConnected
	}
	client.m.Unlock()
	for {
		err := client.UnsafePush(data)
		if err != nil {
			client.errlog.Println("push failed. Retrying...")
			select {
			case <-client.done:
				return errShutdown
			case <-time.After(resendDelay):
			}
			continue
		}
		confirm := <-client.notifyConfirm
		if confirm.Ack {
			client.infolog.Printf("push confirmed [%d]", confirm.DeliveryTag)
			return nil
		}
	}
}

func (client *Client) UnsafePush(data []byte) error {
	client.m.Lock()
	if !client.isReady {
		client.m.Unlock()
		return errNotConnected
	}
	client.m.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return client.channel.PublishWithContext(
		ctx,
		"",               // Exchange
		client.queueName, // Routing key
		false,            // Mandatory
		false,            // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
}

func (client *Client) handleReInit(conn *amqp.Connection) bool {
	for {
		client.m.Lock()
		client.isReady = false
		client.m.Unlock()

		err := client.init(conn)
		if err != nil {
			client.errlog.Println("failed to initialize channel, retrying...")

			select {
			case <-client.done:
				return true
			case <-client.notifyConnClose:
				client.infolog.Println("connection closed, reconnecting...")
				return false
			case <-time.After(reInitDelay):
			}
			continue
		}

		select {
		case <-client.done:
			return true
		case <-client.notifyConnClose:
			client.infolog.Println("connection closed, reconnecting...")
			return false
		case <-client.notifyChanClose:
			client.infolog.Println("channel closed, re-running init...")
		}
	}
}

func (client *Client) Consume() (<-chan amqp.Delivery, error) {
	client.m.Lock()
	if !client.isReady {
		client.m.Unlock()
		return nil, errNotConnected
	}
	client.m.Unlock()

	if err := client.channel.Qos(
		1,     // prefetchCount
		0,     // prefetchSize
		false, // global
	); err != nil {
		return nil, err
	}

	return client.channel.Consume(
		client.queueName,
		"",    // Consumer
		false, // Auto-Ack
		false, // Exclusive
		false, // No-local
		false, // No-Wait
		nil,   // Args
	)
}

func (client *Client) Close() error {
	client.m.Lock()

	defer client.m.Unlock()

	if !client.isReady {
		return errAlreadyClosed
	}
	close(client.done)
	err := client.channel.Close()
	if err != nil {
		return err
	}

	err = client.connection.Close()
	if err != nil {
		return err
	}

	client.isReady = false
	return nil
}
