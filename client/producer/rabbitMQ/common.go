package rabbitmq

import "github.com/streadway/amqp"

const QueueName string = "message"

var (
	serverRabbitMQSvc  string
	serverRabbitMQPort string
)
var Channel *amqp.Channel

var isSend bool = false
