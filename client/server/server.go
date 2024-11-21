package server

import (
	"client/utils"
	"log"
	"net/http"
)

func NewServer() *http.Server {
	muxRouter := newRouter()
	port := utils.GetEnv("Port", "3317")
	log.Println("Running server at http://127.0.0.1:" + port)
	return &http.Server{
		Addr:    ":" + port,
		Handler: muxRouter,
	}
}
