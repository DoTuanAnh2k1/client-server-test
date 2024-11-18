package problem

import (
	"client/common"
	"client/utils"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

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
	listIp := utils.LookupListIp(serverHeadlessSvc)

	// Init tcp connection of client B1
	urlB1 := common.Scheme + listIp[0] + ":" + serverPort + common.PathProblem + "1"
	clientB1.Get(urlB1)

	// Init tcp connection of client B2
	urlB2 := common.Scheme + listIp[1] + ":" + serverPort + common.PathProblem + "2"
	clientB2.Get(urlB2)

	// Init tcp connection of client B3
	urlB3 := common.Scheme + listIp[2] + ":" + serverPort + common.PathProblem + "3"
	clientB3.Get(urlB3)
}

// So now we delete pod B3, B3 will come back soon, during pod B3 down,
// client B3 lost tcp connection to B3. We will use client B3 send a request
// to B through k8s service. We expect that client B3 gonna connect to
// pod B1 or B2.

func ReconnectClient() {
	url := common.Scheme + serverSvc + ":" + serverPort + common.PathProblem + "3"
	clientB3.Get(url)
}

// We are getting B3 up now. We use three client to send request to server B
// through k8s service. If my hypothesis is correct then B3 will not
// get any request, B1 and B2 each will get one requets from client B1
// and client B2, client B3 will send request to pod B1 or B2 depend on
// which pod it have been send when try reconnect

func ProblemDo() {
	urlBase := common.Scheme + serverSvc + ":" + serverPort + common.PathProblem

	clientB1.Get(urlBase + "1")
	clientB2.Get(urlBase + "2")
	clientB3.Get(urlBase + "3")
}
