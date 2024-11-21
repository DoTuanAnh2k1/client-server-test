package main

import (
	"kafka-test/client"
	"kafka-test/server"
)

func main() {
	go client.NewClient()
	server.NewServer()
}
