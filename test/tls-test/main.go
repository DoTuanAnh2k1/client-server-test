package main

import (
	"time"
	"tls-test/client"
	"tls-test/server"
)

func main() {
	go server.NewServer()
	time.Sleep(5 * time.Second)
	client.NewClient()
}
