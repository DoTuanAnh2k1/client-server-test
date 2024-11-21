package kafka

import (
	"client/common"
	"context"
	"log"
	"time"
)

func sendreq() {
	for {
		if !isSend {
			time.Sleep(10 * time.Second)
			continue
		}
		for i := 0; i < common.TicketLength; i++ {
			for j := 0; j < common.Rate/common.TicketLength; j++ {
				go writer.WriteMessages(context.Background(), message)
			}
			time.Sleep(time.Duration(common.Rate/common.TicketLength) * time.Millisecond)
		}
	}
}

func sendOne() {
	err := writer.WriteMessages(context.Background(), message)
	if err != nil {
		log.Printf("Failed to write message: %v", err)
	}
}