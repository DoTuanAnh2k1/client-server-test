package rabbitmq

import (
	"client/utils"

	"github.com/streadway/amqp"
)

func Init() error {
	serverRabbitMQSvc = utils.GetEnv("serverRabbitMQSvc", "localhost")
	serverRabbitMQPort = utils.GetEnv("ServerRabbitMQPort", "5672")
	var err error
	Channel, err = initQueue()
	if err != nil {
		return err
	}
	go sendReq()
	return nil
}

func initQueue() (*amqp.Channel, error) {
	rabbitMQPath := "amqp://guest:guest@" + serverRabbitMQSvc + ":" + serverRabbitMQPort + "/"
	connection, err := amqp.Dial(rabbitMQPath)
	if err != nil {
		return nil, err
	}
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return channel, nil
}
