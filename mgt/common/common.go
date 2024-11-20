package common

import (
	"net"
	"mgt/utils"
)

type PodInfo struct {
	Ip          net.Ip `json:"ip"`
	NumberOfReq uint64 `json:"rate"`
	SuccessRate int `json:"rate-success"`
	InitRequest uint64 `json:"init-req"`
}

type ServerResp struct {
	Request     uint64 `json:"request"`
	InitRequest uint64 `json:"init-req"`
	SuccessRate int `json:"success-rate"`
}

type ServerMeasure struct {
	Request uint64 `json:"request"`
	SucessRate int `json:"success-rate"`
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
	RabbitMQ        = "rabbit-mq"
	Kafka           = "kafka"
	GRPC            = "grpc"
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
