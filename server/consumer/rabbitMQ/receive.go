package rabbitmq

import (
	"log"
	"server/common"
)

func receive() {
	messQueue, err := session.Stream()
	if err != nil {
		log.Println(err)
		// panic(err)
		return
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
