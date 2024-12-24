package server

import (
	"http-gateway/metric"
	"http-gateway/service"
	"net/http"
)

type Router struct {
	Name string 
	Method string
	Pattern string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
}

var (
	metricRouter []Router
	serviceRouter []Router
)

func newMetricRouter() *http.ServeMux {
	mux := http.NewServeMux()
	muxSubPath := http.NewServeMux()
	
	for _, router := range metricRouter {
		muxSubPath.HandleFunc(router.Pattern, router.HandlerFunc)
	}

	mux.Handle("/nsmf-metric/v1/", http.StripPrefix("/nsmf-metric/v1/", muxSubPath))
	
	return mux
}

func newServiceRouter() *http.ServeMux {
	mux := http.NewServeMux()
	muxSubPath := http.NewServeMux()
	
	for _, router := range serviceRouter {
		muxSubPath.HandleFunc(router.Pattern, router.HandlerFunc)
	}

	mux.Handle("/nsmf-svc/v1/", http.StripPrefix("/nsmf-svc/v1/", muxSubPath))
	
	return mux
}

func addHandleMetricRouter() {
	metricRouter = []Router{
		{
			Name: "Heart Beat",
			Method: "GET",
			Pattern: "/heart-beat",
			HandlerFunc: metric.HeartBeatHandler,
		},
	}
}

func addHandleServiceRouter() {
	serviceRouter = []Router {
		{
			Name: "Test",
			Method: "GET",
			Pattern: "/test",
			HandlerFunc: service.TestHandler,
		},
	}
}
