package rabbitmq

import (
	"client/utils"
	"log"
	"time"
)

func Init() {
	serverRabbitMQSvc = utils.GetEnv("ServerRabbitMQSvc", "localhost")
	serverRabbitMQPort = utils.GetEnv("ServerRabbitMQPort", "5672")
	rabbitMQAddr := "amqp://guest:guest@" + serverRabbitMQSvc + ":" + serverRabbitMQPort + "/"
	session = NewRabbitMQSession(queueName, rabbitMQAddr)

	log.Println("Client Rabbit MQ ready")
	time.Sleep(2 * time.Second)
	go sendReq()
}
