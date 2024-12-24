package main

import (
	"http-gateway/server"
)

func main() {
	go server.StartMetricServer()

	server.StartServiceServer()
}