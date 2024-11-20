package rabbitmq

import (
	"client/utils"
)

func Init() {
	serverRabbitMQSvc = utils.GetEnv("ServerRabbitMQSvc", "localhost")
	serverRabbitMQPort = utils.GetEnv("ServerRabbitMQPort", "5672")
	rabbitMQAddr := "amqp://guest:guest@:" + serverRabbitMQSvc + ":" + serverRabbitMQPort + "/"
	session = NewRabbitMQSession(queueName, rabbitMQAddr)
	<-session.notifyReady
	go sendReq()
}
