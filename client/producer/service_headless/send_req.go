package serviceheadless

import (
	"bytes"
	"client/common"
	"net/http"
	"time"
)

func sendReq() {
	body := []byte(common.MessageBody)
	for {
		if !isHeadlessSvc {
			time.Sleep(common.TimeSleep * time.Second)
			continue
		}
		
		for i := 0; i < common.TicketLength * common.Rate / 1000; i ++ {
			go func() {
				connection := getConnection()
				client := getClient(connection.ClientList)
				req, _ := http.NewRequest(http.MethodPost, connection.UrlTest, bytes.NewBuffer(body))
				client.Do(req)
			}()
			time.Sleep(time.Duration(common.TicketLength) * time.Millisecond)
		}
	}
}

func getConnection() *common.Connection {
	indexConnectionGet++
	return connectionList[indexConnectionGet%len(connectionList)]
}

func getClient(clientList []*http.Client) *http.Client {
	indexClientGet++
	return clientList[indexClientGet%common.NumberOfClient]
}
