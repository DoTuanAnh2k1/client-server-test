package server

import (
	"fmt"
	"rabbit-test/session"
	"time"
)

func NewServer(id int) {
	queueName := "queue"
	rabbitMQAddr := "amqp://guest:guest@localhost:5672/"
	session := session.NewRabbitMQSession(queueName, rabbitMQAddr)
	// select {
	// case <-session.NotifyReady:
	// }
	time.Sleep(1 * time.Second)
	fmt.Println("server ready")
	msgQueue, err := session.Stream()
	if err != nil {
		panic(err)
	}
	for msg := range msgQueue {
		fmt.Printf("server %d mess: %s\n", id, string(msg.Body))
		msg.Ack(true)
	}
}
