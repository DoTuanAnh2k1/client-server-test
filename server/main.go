package main

import (
	"server/common"
	rabbitmq "server/consumer/rabbitMQ"
	"server/server"
)

func Init() {
	rabbitmq.Init()
}

func main() {
	server := server.NewServer()
	go common.UpdateInterval()
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
