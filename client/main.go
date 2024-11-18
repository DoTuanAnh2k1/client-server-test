package main

import (
	"client/problem"
	rabbitmq "client/producer/rabbitMQ"
	"client/producer/service"
	serviceheadless "client/producer/service_headless"
	"client/server"
)

func Init() {
	serviceheadless.Init()
	service.Init()
	problem.Init()
	rabbitmq.Init()
}

func main() {
	Init()

	clientServer := server.NewServer()
	err := clientServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
