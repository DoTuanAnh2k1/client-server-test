package kafka

import (
	"context"
	"log"
	"server/common"
)

func receive() {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error while receiving message: %v", err)
			return
		}
		common.CountRequestStart++
		if string(msg.Value) == common.MessageBody {
			common.CountRequestSuccess++
		}

		// log.Printf("received message from kafka on %s: %s\n", msg.Topic, string(msg.Value))
		// log.Printf("Message received: Key=%s, Value=%s, Partition=%d, Offset=%d\n",
		// 	string(msg.Key), string(msg.Value), msg.Partition, msg.Offset)
	}
}
