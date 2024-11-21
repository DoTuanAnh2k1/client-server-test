package main

import (
	"rabbit-test/client"
	"rabbit-test/server"
)

//  docker run -d --hostname my-rabbit --name some-rabbit -p 5672:5672 -p 15672:15672 rabbitmq:3-management

func InitClient() {
	// for i := 1; i < 11; i++ {
	// 	go client.NewClient(i)
	// }
	go client.NewClient(1)
}

func main() {
	InitClient()
	go server.NewServer(1)
	server.NewServer(2)
}
