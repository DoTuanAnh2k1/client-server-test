package rabbitmq

import (
	"server/common"
)

func receive() {
	for msg := range session.Stream() {
		common.CountRequestStart++
		if string(msg.Body) == common.MessageBody {
			common.CountRequestSuccess++
		}
	}
}
