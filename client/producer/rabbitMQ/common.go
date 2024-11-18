package rabbitmq

import "github.com/streadway/amqp"

var QueueName string = "message"

var Channel *amqp.Channel

var isSend bool = false
