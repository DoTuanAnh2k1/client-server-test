package rabbitmq

import (
	"server/common"
)

func receive() {
	messQueue, err := session.Stream()
	if err != nil {
		panic(err)
	}
	for msg := range messQueue {
		common.CountRequestStart++
		if string(msg.Body) == common.MessageBody {
			// log.Println(string(msg.Body))
			common.CountRequestSuccess++
		}
		msg.Ack(true)
	}
}
