package server

import (
	"fmt"
	"net/http"
	"os"
	"server/common"
	"server/consumer/test"
)

func initHandler(w http.ResponseWriter, r *http.Request) {
	common.CountRequestInit++
	fmt.Println("Init Handler #", common.CountRequestInit)
	w.WriteHeader(http.StatusOK)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := os.ReadAll(r.Body)
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

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", testHandler)
	mux.HandleFunc("/init", initHandler)
	mux.HandleFunc("/info", infoHandler)
	mux.HandleFunc("/problem", problemHandler)
	mux.HandleFunc("/measure", measure)
	return mux
}
