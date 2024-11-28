package rabbitmq

import (
	"client/common"
	"log"
	"time"
)

func sendReq() {
	body := []byte(common.MessageBody)
	for {
		if !isSend {
			time.Sleep(common.TimeSleep * time.Second)
			continue
		}
		for i := 0; i < common.TicketLength * common.Rate / 1000; i ++ {
			go session.Push(body)
			time.Sleep(time.Duration(common.TicketLength) * time.Millisecond)
		}
	}
}

func sendOneReq() {
	body := []byte(common.MessageBody)
	err := session.Push(body)
	if err != nil {
		log.Println(err)
	}
}
