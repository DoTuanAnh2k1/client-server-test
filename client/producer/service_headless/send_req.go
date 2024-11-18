package serviceheadless

import (
	"client/common"
	"net/http"
	"time"
)

func sendReq() {
	for {
		if !isHeadlessSvc {
			time.Sleep(10 * time.Second)
			continue
		}
		connection := getConnection()
		client := getClient(connection.ClientList)
		go client.Get(connection.UrlTest)
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
