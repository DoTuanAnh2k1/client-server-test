package rabbitmq

import "github.com/streadway/amqp"

const queueName string = "message"

const (
	// When reconnecting to the server after connection failure
	ReconnectDelay = 5 * time.Second

	// When setting up the channel after a channel exception
	ReInitDelay = 2 * time.Second

	// When resending messages the server didn't confirm
	ResendDelay = 5 * time.Second
)

var (
	serverRabbitMQSvc  string
	serverRabbitMQPort string
)

var (
	errNotConnected  = errors.New("not connected to a server")
	errAlreadyClosed = errors.New("already closed: not connected to the server")
	errShutdown      = errors.New("session is shutting down")
)

var session *Session
