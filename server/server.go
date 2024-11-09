package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var (
	CountRequestStart uint64
	CountRequestPrev  uint64
	CountRequestRate  uint64
	CountRequestInit  uint64
)

type ServerResp struct {
	Request     uint64 `json:"request"`
	InitRequest uint64 `json:"init-req"`
}

func InitHandler(w http.ResponseWriter, r *http.Request) {
	CountRequestInit++
	fmt.Println("Init Handler #", CountRequestInit)
	w.WriteHeader(http.StatusOK)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	CountRequestStart++
	w.WriteHeader(http.StatusOK)
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	var resp ServerResp
	resp.Request = CountRequestRate
	resp.InitRequest = CountRequestInit

	body, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func ProblemHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Println("----------------------------------")
	fmt.Println("This is " + name + "'s connection")
	fmt.Println("----------------------------------")
	w.WriteHeader(http.StatusOK)
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", TestHandler)
	mux.HandleFunc("/init", InitHandler)
	mux.HandleFunc("/info", InfoHandler)
	mux.HandleFunc("/problem", ProblemHandler)
	return mux
}

func newServer() *http.Server {
	h2s := http2.Server{}
	mux := newRouter()
	return &http.Server{
		Addr:    ":3654",
		Handler: h2c.NewHandler(mux, &h2s),
	}
}

func UpdateInterval() {
	for {
		CountRequestRate = CountRequestStart - CountRequestPrev
		CountRequestPrev = CountRequestStart
		time.Sleep(1 * time.Second)
	}
}

func main() {
	server := newServer()
	go UpdateInterval()
	fmt.Println("Running server at http://127.0.0.1:3654")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
