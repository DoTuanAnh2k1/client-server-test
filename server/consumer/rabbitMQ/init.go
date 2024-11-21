package rabbitmq

import (
	"server/utils"
	"time"
)

func Init() {
	serverRabbitMQSvc = utils.GetEnv("ServerRabbitMQSvc", "localhost")
	serverRabbitMQPort = utils.GetEnv("ServerRabbitMQPort", "5672")
	rabbitMQAddr := "amqp://guest:guest@" + serverRabbitMQSvc + ":" + serverRabbitMQPort + "/"
	session = NewRabbitMQSession(queueName, rabbitMQAddr)
	// <-session.notifyReady
	time.Sleep(2 * time.Second)
	go receive()
}
