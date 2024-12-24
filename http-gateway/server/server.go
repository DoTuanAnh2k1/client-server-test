package server

import (
	"net/http"
	"http-gateway/utils"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func StartMetricServer() {
	h2s := http2.Server{}
	mux := newMetricRouter()
	port := utils.GetEnv("MetricPort", "33363")
	
	mertricServer := &http.Server{
		Addr:    ":" + port,
		Handler: h2c.NewHandler(mux, &h2s),
	}

	err := mertricServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func StartServiceServer() {
	h2s := http2.Server{}
	mux := newServiceRouter()
	port := utils.GetEnv("ServicePort", "2194")
	serviceServer := &http.Server{
		Addr:    ":" + port,
		Handler: h2c.NewHandler(mux, &h2s),
	}

	err := serviceServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}