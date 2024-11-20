package main

import (
	"rabbit-test/client"
	"rabbit-test/server"
)

func InitClient() {
	for i := 1; i < 11; i++ {
		go client.NewClient(i)
	}
}

func main() {
	InitClient()
	go server.NewServer(1)
	server.NewServer(2)
}
