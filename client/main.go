package main

import (
	"client/common"
	"client/problem"
	"client/producer/grpc"
	"client/producer/kafka"
	rabbitmq "client/producer/rabbitMQ"
	"client/producer/service"
	serviceheadless "client/producer/service_headless"
	"client/server"
	"client/utils"
	"strconv"
)

// https://www.youtube.com/watch?v=w8xWTIFU4C8
func Init() {
	ticketLength := utils.GetEnv("TicketLength", "500")
	rate := utils.GetEnv("Rate", "10000")
	common.TicketLength, _ = strconv.Atoi(ticketLength)
	common.Rate, _ = strconv.Atoi(rate)
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
