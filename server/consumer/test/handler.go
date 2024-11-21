package test

import (
	"encoding/json"
	"log"
	"server/common"
	"time"
)

func HandlerTest(body string) {
	common.CountRequestStart++
	if body == common.MessageBody {
		common.CountRequestSuccess++
	}
}

func HandlerInfo() ([]byte, error) {
	var resp common.ServerResp
	resp.Request = common.CountRequestRate
	resp.InitRequest = common.CountRequestInit

	body, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func HandlerProblem(name string) {
	log.Println("---------------------------------")
	log.Println("This is " + name + "'s connection")
	log.Println("---------------------------------")
}

func HandlerMeasure() ([]byte, error) {
	ans := uint64(0)
	for i := 0; i < 10; i++ {
		ans = ans + common.CountRequestRate/10
		time.Sleep(1 * time.Second)
	}
	var serverMeasure common.ServerMeasure
	serverMeasure.Request = ans

	body, err := json.Marshal(serverMeasure)
	if err != nil {
		return nil, err
	}
	return body, nil
}
