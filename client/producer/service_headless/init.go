package serviceheadless

import (
	"client/common"
	"client/utils"
	"context"
	"crypto/tls"
	"net"
	"net/http"

	"golang.org/x/net/http2"
)

func Init() {
	serverHeadlessSvc = utils.GetEnv("ServerHeadlessSvc", "localhost")
	serverPort = utils.GetEnv("ServerPort", "3654")
	initConnectionList()
	go ScanIp()
	go sendReq()
}

func initConnectionList() {
	// Init list connection, each connection have NumberOfClient clients
	ipList := utils.LookupListIp(serverHeadlessSvc)
	for _, ip := range ipList {
		connection := InitConnection(ip, serverPort)
		connectionList = append(connectionList, connection)
	}
}

func InitConnection(ip, serverPort string) *common.Connection {
	url := common.Scheme + ip + ":" + serverPort + common.PathInit
	connection := &common.Connection{
		Ip:         ip,
		UrlTest:    common.Scheme + ip + ":" + serverPort + common.PathTest,
		ClientList: make([]*http.Client, common.NumberOfClient),
	}
	for i := 0; i < common.NumberOfClient; i++ {
		client := &http.Client{
			Transport: &http2.Transport{
				AllowHTTP:          true,
				DisableCompression: true,
				DialTLSContext:     dialTlsContext,
			},
		}
		// init tcp connection
		go client.Get(url)
		connection.ClientList[i] = client
	}
	return connection
}

func dialTlsContext(ctx context.Context, network, address string, cfg *tls.Config) (net.Conn, error) {
	return net.Dial(network, address)
}
