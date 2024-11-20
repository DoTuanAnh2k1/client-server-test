package common

import "net/http"

const (
	NumberOfClient = 8
	Scheme         = "http://"
	PathInit       = "/init"
	PathTest       = "/test"
	PathProblem    = "/problem?name=client"
)

const (
	TimeSleep = 10
)

const MessageBody = "HeHe"

type Connection struct {
	Ip         string
	UrlTest    string
	ClientList []*http.Client
}

var TickLength int
