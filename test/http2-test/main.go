package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	HTTPVersion1 = 1
	HTTPVersion2 = 2
)

func NewHTTPServer(httpServerVersion uint8) {
	var s *http.Server
	switch httpServerVersion {
	case HTTPVersion1:
		s = newServerVersion1()
	case HTTPVersion2:
		s = newServerVersion2()
	default:
		panic("wrong version")
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func newServerVersion1() *http.Server {
	mux := newRouter()
	return &http.Server{
		Addr:    ":8457",
		Handler: mux,
	}
}

func newServerVersion2() *http.Server {
	h2s := &http2.Server{}
	mux := newRouter()
	return &http.Server{
		Addr:    ":8457",
		Handler: h2c.NewHandler(mux, h2s),
	}
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", reqHandler)
	return mux
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func NewClientVer2() {
	c := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP:          true,
			DisableCompression: true,
			DialTLSContext:     dial,
		},
		Timeout: 2 * time.Second,
	}
	resp, err := c.Get("http://localhost:8457/")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Body)
}

func NewClientVer1() {
	c := &http.Client{}
	resp, err := c.Get("http://localhost:8457/")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Body)
}

func dial(ctx context.Context, network, addr string, cfg *tls.Config) (net.Conn, error) {
	return net.Dial(network, addr)
}

func main() {
	go NewHTTPServer(HTTPVersion2)
	// go NewHTTPServer(HTTPVersion1)
	// NewClientVer2()
	NewClientVer1()
}
