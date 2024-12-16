package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewHTTPServer() {
	s := newServerVersion()
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func newServerVersion() *http.Server {
	h2s := &http2.Server{}
	mux := newRouter()
	return &http.Server{
		Addr:    ":8457",
		Handler: h2c.NewHandler(mux, h2s),
	}
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/post", postHandler)
	return mux
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler post request")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func NewClientVer2() {
	c := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP:          true,
			DisableCompression: true,
		},
		Timeout: 2 * time.Second,
	}
	bodyMess := "Abc-ASDJIOEF"
	url := "http://localhost:8457/post"
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(bodyMess)))
	req.Header.Set("Abc-ASDJIOEF-----b", "value")
	// resp, err := c.Get("http://localhost:8457/")
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Body)
}


func main() {
	NewHTTPServer()
}
