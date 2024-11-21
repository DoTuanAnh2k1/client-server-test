package main

import (
	"server/common"
	"server/consumer/kafka"
	rabbitmq "server/consumer/rabbitMQ"
	"server/server"
)

func Init() {
	rabbitmq.Init()
	kafka.Init()
}

func main() {
	go Init()
	server := server.NewServer()
	go common.UpdateInterval()
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
