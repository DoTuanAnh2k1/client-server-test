package kafka

import (
	"context"
	"log"
)

func receive() {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Error while receiving message: %v", err)
		}

		log.Printf("received message from kafka on %s: %s\n", msg.Topic, string(msg.Value))
		// log.Printf("Message received: Key=%s, Value=%s, Partition=%d, Offset=%d\n",
		// 	string(msg.Key), string(msg.Value), msg.Partition, msg.Offset)
	}
}
