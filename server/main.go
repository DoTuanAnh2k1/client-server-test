package main

import (
	"server/common"
	"server/server"
)

func main() {
	server := server.NewServer()
	go common.UpdateInterval()
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
