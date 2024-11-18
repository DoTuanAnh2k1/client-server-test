package service

import (
	"client/common"
	"net/http"
	"time"
)

func sendReq() {
	url := common.Scheme + serverSvc + ":" + serverPort + common.PathTest
	for {
		if !isSvc {
			time.Sleep(common.TimeSleep * time.Second)
			continue
		}
		client := getClient(clientList)
		go client.Get(url)
	}
}

func getClient(clientList []*http.Client) *http.Client {
	indexClientGet++
	return clientList[indexClientGet%common.NumberOfClient]
}
