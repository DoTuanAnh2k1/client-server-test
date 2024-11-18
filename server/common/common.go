package common

import "time"

var (
	CountRequestStart uint64
	CountRequestPrev  uint64
	CountRequestRate  uint64
	CountRequestInit  uint64
)

type ServerResp struct {
	Request     uint64 `json:"request"`
	InitRequest uint64 `json:"init-req"`
}

type ServerMeasure struct {
	Request uint64 `json:"request"`
}

const MessageBody = "HeHe"

func UpdateInterval() {
	for {
		CountRequestRate = CountRequestStart - CountRequestPrev
		CountRequestPrev = CountRequestStart
		time.Sleep(1 * time.Second)
	}
}
