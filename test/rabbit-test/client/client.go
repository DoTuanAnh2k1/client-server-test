package client

import (
	"fmt"
	"rabbit-test/session"
	"time"
)

func NewClient(id int) {
	queueName := "queue"
	rabbitMQAddr := "amqp://guest:guest@localhost:5672/"
	message := "message"
	session := session.NewRabbitMQSession(queueName, rabbitMQAddr)
	// select {
	// case <-session.NotifyReady:
	// }
	fmt.Println("client ready")
	i := 0
	for {
		time.Sleep(1 * time.Second)
		tmp := []byte(fmt.Sprintf("%s %d from client %d", message, i, id))
		i++
		err := session.Push(tmp)
		if err != nil {
			panic(err)
		}
	}
}
