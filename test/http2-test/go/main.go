package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"sync"
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
	mux.HandleFunc("/5", reqDelay5sHandler)
	mux.HandleFunc("/10", reqDelay10sHandler)
	mux.HandleFunc("/15", reqDelay15sHandler)
	return mux
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func reqDelay5sHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("5 second delay")
	time.Sleep(5 * time.Second)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func reqDelay10sHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("10 second delay")
	time.Sleep(10 * time.Second)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func reqDelay15sHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("15 second delay")
	time.Sleep(15 * time.Second)
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

func multiflexingTest() {
	c1 := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP:          true,
			DisableCompression: true,
			DialTLSContext:     dial,
		},
		Timeout: 2 * time.Second,
	}
	go NewHTTPServer(HTTPVersion2)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		c1.Get("http://localhost:8457/15")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		c1.Get("http://localhost:8457/10")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		c1.Get("http://localhost:8457/5")
		wg.Done()
	}()

	wg.Wait()
}

func main() {
	// go NewHTTPServer(HTTPVersion2)
	// // go NewHTTPServer(HTTPVersion1)
	// // NewClientVer2()
	// NewClientVer1()
	multiflexingTest()
}
