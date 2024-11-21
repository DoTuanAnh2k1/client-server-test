package server

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func NewServer() {
	// Địa chỉ Kafka broker và topic
	brokerAddress := "localhost:9092"
	topic := "example-topic"
	groupID := "example-group"

	// Tạo Kafka reader (consumer)
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerAddress},
		GroupID:  groupID, // Consumer group ID
		Topic:    topic,   // Topic cần lắng nghe
		MinBytes: 10e3,    // Dữ liệu tối thiểu nhận được (10KB)
		MaxBytes: 10e6,    // Dữ liệu tối đa nhận được (10MB)
	})

	defer reader.Close()

	log.Println("Consumer started, listening for messages...")

	// Lắng nghe và xử lý message
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Error while receiving message: %v", err)
		}

		// In nội dung message
		log.Printf("Message received: Key=%s, Value=%s, Partition=%d, Offset=%d\n",
			string(msg.Key), string(msg.Value), msg.Partition, msg.Offset)
	}
}
