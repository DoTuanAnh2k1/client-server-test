package main

import (
	"context"
	"crypto/tls"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func IntRandom(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

type Connection struct {
	Ip         string
	UrlTest    string
	ClientList []*http.Client
}

var ConnectionList []*Connection

const NumberOfClient = 8

func dialTlsContext(ctx context.Context, network, address string, cfg *tls.Config) (net.Conn, error) {
	return net.Dial(network, address)
}

func lookupListIp(headlessSvc string) []string {
	ips, err := net.LookupIP(headlessSvc)
	if err != nil {
		panic(err)
	}
	var listIp []string
	for _, v := range ips {
		listIp = append(listIp, v.String())
	}
	return listIp
}

func InitConnection(headlessSvc string, serverPort string) {
	ipList := lookupListIp(headlessSvc)
	for _, ip := range ipList {
		url := "http://" + ip + ":" + serverPort + "/init"
		connection := &Connection{
			Ip:         ip,
			UrlTest:    "http://" + ip + ":" + serverPort + "/test",
			ClientList: make([]*http.Client, NumberOfClient),
		}
		for i := 0; i < NumberOfClient; i++ {
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
		ConnectionList = append(ConnectionList, connection)
	}
}

func getClient(clientList []*http.Client) *http.Client {
	return clientList[IntRandom(0, NumberOfClient-1)]
}

func sendReq() {
	for {
		for _, connection := range ConnectionList {
			client := getClient(connection.ClientList)
			client.Get(connection.UrlTest)
		}
	}
}

func newGin() *gin.Engine {
	r := gin.Default()
	r.GET("/send-req", func(c *gin.Context) {
		go sendReq()
		c.String(200, "ok")
	})
	return r
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}
	return val
}

func main() {
	headlessSvc := getEnv("ServerHeadlessSvc", "localhost")
	serverPort := getEnv("ServerPort", "3654")
	InitConnection(headlessSvc, serverPort)
	clientServer := newGin()
	err := clientServer.Run(":3317")
	if err != nil {
		panic(err)
	}
}
