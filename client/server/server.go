package server

import (
	"client/utils"
	"net/http"
)

func NewServer() *http.Server {
	muxRouter := newRouter()
	port := utils.GetEnv("PORT", "3654")
	return &http.Server{
		Addr:    ":" + port,
		Handler: muxRouter,
	}
}
