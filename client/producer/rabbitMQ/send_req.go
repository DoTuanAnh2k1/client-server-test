package rabbitmq

import (
	"client/common"
	"fmt"
	"time"
)

func sendReq() {
	body := []byte(common.MessageBody)
	for {
		if !isSend {
			time.Sleep(common.TimeSleep * time.Second)
			continue
		}
		session.Push(body)
	}
}

func sendOneReq() {
	body := []byte(common.MessageBody)
	err := session.Push(body)
	if err != nil {
		fmt.Println(err)
	}
}
