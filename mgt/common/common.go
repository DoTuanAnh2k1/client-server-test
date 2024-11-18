package common

import "mgt/utils"

type PodInfo struct {
	Ip          string `json:"ip"`
	NumberOfReq uint64 `json:"number-of-request"`
	InitRequest uint64 `json:"init-req"`
}

type ServerResp struct {
	Request     uint64 `json:"request"`
	InitRequest uint64 `json:"init-req"`
}

type ServerMeasure struct {
	Request uint64 `json:"request"`
}

var (
	ClientSvc         string
	ClientPort        string
	ServerHeadlessSvc string
	ServerPort        string
)

const (
	Service         = "svc"
	ServiceHeadless = "svc-headless"
)

const (
	ProblemInit      = "init"
	ProblemReconnect = "reconnect"
	ProblemDo        = "do"
)

func InitVar() {
	ClientSvc = utils.GetEnv("ClientSvc", "localhost")
	ClientPort = utils.GetEnv("ClientPort", "3317")
	ServerHeadlessSvc = utils.GetEnv("ServerHeadlessSvc", "localhost")
	ServerPort = utils.GetEnv("ServerPort", "3654")
}
