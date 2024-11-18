package server

import (
	"fmt"
	"net/http"
	"server/utils"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewServer() *http.Server {
	port := utils.GetEnv("Port", "3654")
	fmt.Println("Running server at http://127.0.0.1:" + port)
	h2s := http2.Server{}
	mux := newRouter()
	return &http.Server{
		Addr:    ":" + port,
		Handler: h2c.NewHandler(mux, &h2s),
	}
}
