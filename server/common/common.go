package common

import "time"

var (
	CountRequestStart uint64
	CountRequestPrev  uint64
	CountRequestRate  uint64
	CountRequestInit  uint64
	CountRequestSuccess uint64
	CountRequestSuccessPrev  uint64
	CountRequestSuccessRate  uint64
)

type ServerResp struct {
	Request     uint64 `json:"request"`
	InitRequest uint64 `json:"init-req"`
	SuccessRate int `json:"success-rate"`
}

type ServerMeasure struct {
	Request uint64 `json:"request"`
	SucessRate int `json:"success-rate"`
}

const MessageBody = "HeHe"

func UpdateInterval() {
	for {
		CountRequestRate = CountRequestStart - CountRequestPrev
		CountRequestPrev = CountRequestStart
		CountRequestSuccessRate = CountRequestSuccess - CountRequestSuccessPrev
		CountRequestSuccessPrev = CountRequestSuccess
		time.Sleep(1 * time.Second)
	}
}
