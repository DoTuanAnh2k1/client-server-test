package problem

import (
	"client/utils"
	"context"
	"crypto/tls"
	"net"
)

var (
	serverHeadlessSvc string
	serverSvc         string
	serverPort        string
)

func Init() {
	serverHeadlessSvc = utils.GetEnv("ServerHeadlessSvc", "localhost")
	serverSvc = utils.GetEnv("ServerPort", "3654")
	serverPort = utils.GetEnv("ServerPort", "3654")
}

func dialTlsContext(ctx context.Context, network, address string, cfg *tls.Config) (net.Conn, error) {
	return net.Dial(network, address)
}
