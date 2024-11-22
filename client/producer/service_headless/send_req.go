package serviceheadless

import (
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
		connection := getConnection()
		client := getClient(connection.ClientList)
		for i := 0; i < common.TicketLength; i++ {
			for j := 0; j < common.Rate / common.TicketLength; j++ {
				go func() {
					req, _ := http.NewRequest(http.MethodPost, connection.UrlTest, bytes.NewBuffer(body))
					client.Do(req)
				}()
			}
			time.Sleep(time.Duration(common.Rate / common.TicketLength) * time.Millisecond)
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
