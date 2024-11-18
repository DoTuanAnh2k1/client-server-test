package service

import (
	"client/common"
	"client/utils"
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func Init() {
	serverSvc = utils.GetEnv("ServerPort", "3654")
	serverPort = utils.GetEnv("ServerSvc", "localhost")
	initClientList()
	go sendReq()
}

func initClientList() {
	url := common.Scheme + serverSvc + ":" + serverPort + common.PathInit
	for i := 0; i < common.NumberOfClient; i++ {
		client := &http.Client{
			Transport: &http2.Transport{
				AllowHTTP:          true,
				DisableCompression: true,
				DialTLSContext:     dialTlsContext,
			},
			Timeout: 2 * time.Second,
		}
		go client.Get(url)
		clientList = append(clientList, client)
	}
}

func dialTlsContext(ctx context.Context, network, address string, cfg *tls.Config) (net.Conn, error) {
	return net.Dial(network, address)
}
