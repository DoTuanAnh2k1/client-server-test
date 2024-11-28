package service

import (
	"bytes"
	"client/common"
	"net/http"
	"time"
)

func sendReq() {
	url := common.Scheme + serverSvc + ":" + serverPort + common.PathTest
	body := []byte(common.MessageBody)

	for {
		if !isSvc {
			time.Sleep(common.TimeSleep * time.Second)
			continue
		}
		client := getClient(clientList)
		for i := 0; i < common.TicketLength * common.Rate / 1000; i ++ {
			go func() {
				req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
				client.Do(req)
			}()
			time.Sleep(time.Duration(common.TicketLength) * time.Millisecond)
		}
	}
}

func getClient(clientList []*http.Client) *http.Client {
	indexClientGet++
	return clientList[indexClientGet%common.NumberOfClient]
}
