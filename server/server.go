package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var CountRequest uint64
var CountRequestInit uint64

func newGin() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		CountRequest++
		c.String(200, "ok")
	})
	r.GET("/init", func(c *gin.Context) {
		CountRequestInit++
		c.String(200, "ok")
	})
	r.GET("/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"request":  CountRequest,
			"init-req": CountRequestInit,
		})
	})

	return r
}

func newServer() *http.Server {
	h2s := http2.Server{}
	router := newGin()
	return &http.Server{
		Addr:    ":3654",
		Handler: h2c.NewHandler(router, &h2s),
	}
}

func main() {
	server := newServer()
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
