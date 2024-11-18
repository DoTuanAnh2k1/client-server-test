package common

import "net/http"

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
