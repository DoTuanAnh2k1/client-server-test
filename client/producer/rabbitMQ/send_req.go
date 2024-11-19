package rabbitmq

import (
	"client/common"
	"time"

	"github.com/streadway/amqp"
)

func sendReq() {
	body := []byte(common.MessageBody)
	for {
		if !isSend {
			time.Sleep(common.TimeSleep * time.Second)
			continue
		}
		Channel.Publish(
			"",        // exchange
			QueueName, // key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			},
		)
	}
}

func sendOneReq() {
	body := []byte(common.MessageBody)
	Channel.Publish(
		"",        // exchange
		QueueName, // key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}
