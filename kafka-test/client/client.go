package client

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func NewClient() {
	// Địa chỉ Kafka broker và topic
	brokerAddress := "localhost:9092"
	topic := "example-topic"

	// Tạo Kafka writer (producer)
	writer := kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne,
	}

	defer writer.Close()

	// Message cần gửi
	message := kafka.Message{
		Key:   []byte("Key-A"),                   // Key định tuyến (optional)
		Value: []byte("Hello, Kafka Segmentio!"), // Giá trị message
	}

	for {
		time.Sleep(2 * time.Second)
		// Gửi message
		err := writer.WriteMessages(context.Background(), message)
		if err != nil {
			log.Fatalf("Failed to write message: %v", err)
		}

		log.Println("Message sent successfully")
	}
}
