package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/common"
	"time"
)

func initHandler(w http.ResponseWriter, r *http.Request) {
	common.CountRequestInit++
	fmt.Println("Init Handler #", common.CountRequestInit)
	w.WriteHeader(http.StatusOK)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	common.CountRequestStart++
	w.WriteHeader(http.StatusOK)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	var resp common.ServerResp
	resp.Request = common.CountRequestRate
	resp.InitRequest = common.CountRequestInit

	body, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func problemHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Println("---------------------------------")
	fmt.Println("This is " + name + "'s connection")
	fmt.Println("---------------------------------")
	w.WriteHeader(http.StatusOK)
}

func measure(w http.ResponseWriter, r *http.Request) {
	ans := uint64(0)
	for i := 0; i < 10; i++ {
		ans = ans + common.CountRequestRate/10
		time.Sleep(1 * time.Second)
	}
	var serverMeasure common.ServerMeasure
	serverMeasure.Request = ans

	body, err := json.Marshal(serverMeasure)
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
