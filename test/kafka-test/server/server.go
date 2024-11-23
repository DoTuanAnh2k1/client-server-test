package server

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func NewServer(id int) {
	brokerAddress := "localhost:9092"
	topic := "example-topic"
	groupID := "example-group"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerAddress},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
		MaxWait:  500 * time.Millisecond,
	})

	defer reader.Close()

	log.Println("Consumer started, listening for messages...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Error while receiving message: %v", err)
		}

		log.Printf("Server %d received message on %s: %s\n", id, msg.Topic, string(msg.Value))
		// log.Printf("Message received: Key=%s, Value=%s, Partition=%d, Offset=%d\n",
		// 	string(msg.Key), string(msg.Value), msg.Partition, msg.Offset)
	}
}
