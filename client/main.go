package main

import (
	"client/problem"
	"client/producer/grpc"
	"client/producer/kafka"
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
	grpc.Init()
	kafka.Init()
}

func main() {
	go Init()

	clientServer := server.NewServer()
	err := clientServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
