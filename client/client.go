package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
)

const (
	NumberOfClient = 8
	Scheme         = "http://"
	PathInit       = "/init"
	PathTest       = "/test"
	PathProblem    = "/problem?name=client"
)

type Connection struct {
	Ip         string
	UrlTest    string
	ClientList []*http.Client
}

var (
	ConnectionSolutionList []*Connection
	ConnectionList         []*Connection
	ClientList             []*http.Client
)

var (
	ServerHeadlessSvc string
	ServerPort        string
	ServerSvc         string
)

var (
	indexConnectionSolutionGet int
	indexConnectionGet         int
	indexClientGet             int
)

var (
	IsHeadlessSvc   bool = false
	IsSvc           bool = false
	IsSvcSolution   bool = false
	IsSvcInitClient bool = false
)

func dialTlsContextTimeOut(ctx context.Context, network, address string, cfg *tls.Config) (net.Conn, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	go func() {
		time.Sleep(5 * time.Second)
		conn.Close()
	}()

	return conn, nil
}

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

func InitSolutionConnection() *Connection {
	connection := &Connection{
		UrlTest:    Scheme + ServerSvc + ":" + ServerPort + PathTest,
		ClientList: make([]*http.Client, NumberOfClient),
	}
	for i := 0; i < NumberOfClient; i++ {
		connection.ClientList[i] = &http.Client{
			Transport: &http2.Transport{
				AllowHTTP:          true,
				DisableCompression: true,
				DialTLSContext:     dialTlsContextTimeOut,
			},
			Timeout: 2 * time.Second,
		}
	}
	return connection
}

func InitConnection(ip, serverPort string) *Connection {
	url := Scheme + ip + ":" + serverPort + PathInit
	connection := &Connection{
		Ip:         ip,
		UrlTest:    Scheme + ip + ":" + serverPort + PathTest,
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
	return connection
}

func InitConnectionList() {
	// Init list connection, each connection have NumberOfClient clients
	ipList := lookupListIp(ServerHeadlessSvc)
	for _, ip := range ipList {
		connection := InitConnection(ip, ServerPort)
		ConnectionList = append(ConnectionList, connection)
	}

	// Init list client for service
	url := Scheme + ServerSvc + ":" + ServerPort + PathInit
	for i := 0; i < NumberOfClient; i++ {
		client := &http.Client{
			Transport: &http2.Transport{
				AllowHTTP:          true,
				DisableCompression: true,
				DialTLSContext:     dialTlsContext,
			},
			Timeout: 2 * time.Second,
		}
		go client.Get(url)
		ClientList = append(ClientList, client)
	}
}

func InitConnectionSolutionList() {
	ipList := lookupListIp(ServerHeadlessSvc)
	for i := 0; i < len(ipList); i++ {
		connection := InitSolutionConnection()
		ConnectionSolutionList = append(ConnectionSolutionList, connection)
	}
}

func scanIp() {
	ipList := lookupListIp(ServerHeadlessSvc)
	// check remove
	for i := 0; i < len(ConnectionList); i++ {
		connection := ConnectionList[i]
		isExist := false
		for _, ip := range ipList {
			if connection.Ip == ip {
				isExist = true
				break
			}
		}
		if !isExist {
			fmt.Println("Remove Connection: ", connection.Ip)
			ConnectionList = append(ConnectionList[:i], ConnectionList[i+1:]...)
			fmt.Println("Connection List: ", ConnectionList)
			i--
		}
	}

	// check add new
	for _, ip := range ipList {
		isActive := false
		for _, connection := range ConnectionList {
			if connection.Ip == ip {
				isActive = true
				break
			}
		}
		if !isActive {
			fmt.Println("Add Connection: ", ip)
			connection := InitConnection(ip, ServerPort)
			ConnectionList = append(ConnectionList, connection)
			fmt.Println("Connection List: ", ConnectionList)
		}
	}
}

func ScanIp() {
	for {
		time.Sleep(5 * time.Second)
		go scanIp()
	}
}

func getClient(clientList []*http.Client) *http.Client {
	indexClientGet++
	return clientList[indexClientGet%NumberOfClient]
}

func getConnection() *Connection {
	indexConnectionGet++
	return ConnectionList[indexConnectionGet%len(ConnectionList)]
}

func getSolutionConnection() *Connection {
	indexConnectionSolutionGet++
	return ConnectionSolutionList[indexConnectionSolutionGet%len(ConnectionList)]
}

func sendReq() {
	for {
		if !IsHeadlessSvc {
			continue
		}
		connection := getConnection()
		client := getClient(connection.ClientList)
		client.Get(connection.UrlTest)
	}
}

func sendReqService() {
	url := Scheme + ServerSvc + ":" + ServerPort + PathTest
	for {
		if !IsSvc {
			continue
		}
		client := getClient(ClientList)
		client.Get(url)
	}
}

func sendReqServiceInitNewClient() {
	url := Scheme + ServerSvc + ":" + ServerPort + PathTest
	for {
		if !IsSvcInitClient {
			continue
		}
		client := &http.Client{
			Transport: &http2.Transport{
				AllowHTTP:          true,
				DisableCompression: true,
				DialTLSContext:     dialTlsContext,
			},
			Timeout: 2 * time.Second,
		}
		client.Get(url)
	}
}

func sendReqServiceSolution() {
	url := Scheme + ServerSvc + ":" + ServerPort + PathTest
	for {
		if !IsSvcSolution {
			continue
		}
		getClient(getSolutionConnection().ClientList).Get(url)
	}
}

// Problem: imagine we have 3 http2 clients in server client A, each client
// connect to one server pod (B1, B2, B3). These clients are the same.

var (
	clientB1 = &http.Client{
		Transport: &http2.Transport{
			AllowHTTP:          true,
			DisableCompression: true,
			DialTLSContext:     dialTlsContext,
		},
		Timeout: 2 * time.Second,
	}
	clientB2 = &http.Client{
		Transport: &http2.Transport{
			AllowHTTP:          true,
			DisableCompression: true,
			DialTLSContext:     dialTlsContext,
		},
		Timeout: 2 * time.Second,
	}
	clientB3 = &http.Client{
		Transport: &http2.Transport{
			AllowHTTP:          true,
			DisableCompression: true,
			DialTLSContext:     dialTlsContext,
		},
		Timeout: 2 * time.Second,
	}
)

// We init this client's connection by using that client send one http
// request direct to Bi (with i in range 1 to 3). So we have client B1
// connect to B1, client B2 connect to B2, client B3 connect to B3

func InitConnectionForProblem() {
	// Get list ip of server B
	// let's say we have 3 pods only
	listIp := lookupListIp(ServerHeadlessSvc)

	// Init tcp connection of client B1
	urlB1 := Scheme + listIp[0] + ":" + ServerPort + PathProblem + "1"
	clientB1.Get(urlB1)

	// Init tcp connection of client B2
	urlB2 := Scheme + listIp[1] + ":" + ServerPort + PathProblem + "2"
	clientB2.Get(urlB2)

	// Init tcp connection of client B3
	urlB3 := Scheme + listIp[2] + ":" + ServerPort + PathProblem + "3"
	clientB3.Get(urlB3)
}

// So now we delete pod B3, B3 will come back soon, during pod B3 down,
// client B3 lost tcp connection to B3. We will use client B3 send a request
// to B through k8s service. We expect that client B3 gonna connect to
// pod B1 or B2.

func ReconnectClient() {
	url := Scheme + ServerSvc + ":" + ServerPort + PathProblem + "3"
	clientB3.Get(url)
}

// We are getting B3 up now. We use three client to send request to server B
// through k8s service. If my hypothesis is correct then B3 will not
// get any request, B1 and B2 each will get one requets from client B1
// and client B2, client B3 will send request to pod B1 or B2 depend on
// which pod it have been send when try reconnect

func ProblemDo() {
	urlBase := Scheme + ServerSvc + ":" + ServerPort + PathProblem

	clientB1.Get(urlBase + "1")
	clientB2.Get(urlBase + "2")
	clientB3.Get(urlBase + "3")
}

func newGin() *gin.Engine {
	r := gin.Default()

	// Headless svc
	r.GET("/send-req", func(c *gin.Context) {
		IsHeadlessSvc = true
		c.String(http.StatusOK, "ok")
	})
	r.GET("/off-send-req", func(c *gin.Context) {
		IsHeadlessSvc = false
		c.String(http.StatusOK, "ok")
	})

	// Svc
	r.GET("/send-req-svc", func(c *gin.Context) {
		IsSvc = true
		c.String(http.StatusOK, "ok")
	})
	r.GET("/off-send-req-svc", func(c *gin.Context) {
		IsSvc = false
		c.String(http.StatusOK, "ok")
	})

	// Svc with init new client every time send
	r.GET("/send-req-svc-init-client", func(c *gin.Context) {
		IsSvcInitClient = true
		c.String(http.StatusOK, "ok")
	})
	r.GET("/off-send-req-svc-init-client", func(c *gin.Context) {
		IsSvcInitClient = false
		c.String(http.StatusOK, "ok")
	})

	// Problem
	r.GET("/problem-init", func(c *gin.Context) {
		InitConnectionForProblem()
		c.String(http.StatusOK, "ok")
	})
	r.GET("/problem-reconnect", func(c *gin.Context) {
		ReconnectClient()
		c.String(http.StatusOK, "ok")
	})
	r.GET("/problem-do", func(c *gin.Context) {
		ProblemDo()
		c.String(http.StatusOK, "ok")
	})

	// Solution
	r.GET("/send-req-svc-sol", func(c *gin.Context) {
		IsSvcSolution = true
		c.String(http.StatusOK, "ok")
	})
	r.GET("/off-send-req-svc-sol", func(c *gin.Context) {
		IsSvcSolution = false
		c.String(http.StatusOK, "ok")
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

func InitVariable() {
	ServerHeadlessSvc = getEnv("ServerHeadlessSvc", "localhost")
	ServerPort = getEnv("ServerPort", "3654")
	ServerSvc = getEnv("ServerSvc", "localhost")
}

func StartSendReq() {
	go sendReq()
	go sendReqService()
	go sendReqServiceInitNewClient()
	go sendReqServiceSolution()
}

func Init() {
	InitVariable()
	InitConnectionList()
	InitConnectionSolutionList()
}

func main() {
	Init()
	go ScanIp()
	go StartSendReq()
	clientServer := newGin()
	err := clientServer.Run(":3317")
	if err != nil {
		panic(err)
	}
}
