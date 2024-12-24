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
	prefix := "/nsmf-metric/v1"
	mux := http.NewServeMux()
	for _, router := range metricRouter {
		mux.HandleFunc(prefix+router.Pattern, router.HandlerFunc)
	}
	return mux
}

func newServiceRouter() *http.ServeMux {
	prefix := "/nsmf-svc/v1"
	mux := http.NewServeMux()
	for _, router := range serviceRouter {
		mux.HandleFunc(prefix+router.Pattern, router.HandlerFunc)
	}
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
