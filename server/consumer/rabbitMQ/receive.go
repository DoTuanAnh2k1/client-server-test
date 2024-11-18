package rabbitmq

import (
	"server/common"
)

func receive() {
	msgs, err := Channel.Consume(
		QueueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		panic(err)
	}
	for msg := range msgs {
		if string(msg.Body) == common.MessageBody {
			common.CountRequestStart++

		}
	}
}
