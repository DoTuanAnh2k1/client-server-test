package rabbitmq

import "github.com/streadway/amqp"

func Init() {
	var err error
	Channel, err = initQueue()
	if err != nil {
		panic(err)
	}
	go receive()
}

func initQueue() (*amqp.Channel, error) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return channel, nil
}
