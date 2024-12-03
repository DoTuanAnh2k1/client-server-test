package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"server/common"
	"server/consumer/test"
	"server/metric"
)

func initHandler(w http.ResponseWriter, r *http.Request) {
	common.CountRequestInit++
	log.Println("Init Handler #", common.CountRequestInit)
	w.WriteHeader(http.StatusOK)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	go test.HandlerTest(string(body))
	w.WriteHeader(http.StatusOK)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := test.HandlerInfo()
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func problemHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	test.HandlerProblem(name)
	w.WriteHeader(http.StatusOK)
}

func measure(w http.ResponseWriter, r *http.Request) {
	body, err := test.HandlerMeasure()
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func metrics(w http.ResponseWriter, r *http.Request) {
	podMetricInfo, err := metric.GetMetric()
	if err != nil {
		panic(err)
	}
	bodyResp, err := json.Marshal(podMetricInfo)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bodyResp)
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", testHandler)
	mux.HandleFunc("/init", initHandler)
	mux.HandleFunc("/info", infoHandler)
	mux.HandleFunc("/problem", problemHandler)
	mux.HandleFunc("/measure", measure)
	mux.HandleFunc("/metrics", metrics)
	return mux
}
