package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func NewClient(id int) {
	brokerAddress := "localhost:9092"
	topic := "example-topic"

	writer := kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne,
	}
	
	defer writer.Close()

	message := kafka.Message{
		Key:   []byte(fmt.Sprintf("key-%d", id)),                  
		Value: []byte("Hello, Kafka Segmentio!"), 
	}

	log.Printf("Client %d ready for send mess to kafka", id)
	i := 0
	for {
		i++
		time.Sleep(500 * time.Millisecond)
		message.Value = []byte(fmt.Sprintf("Message %d from client %d", i, id) )
		err := writer.WriteMessages(context.Background(), message)
		if err != nil {
			log.Fatalf("Failed to write message: %v", err)
		}

		log.Println("Message sent successfully")
	}
}
