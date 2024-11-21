package grpc

import (
	"client/common"
	pb "client/producer/grpc/proto"
	"context"
	"log"
	"time"
)

func sendReq() {
	for {
		if !isSend {
			time.Sleep(10 * time.Second)
			continue
		}
		for i := 0; i < common.TicketLength; i++ {
			for j := 0; j < common.Rate/common.TicketLength; j++ {
				go clientGRPC.SayHello(context.Background(), &pb.HelloRequest{Message: common.MessageBody})
			}
			time.Sleep(time.Duration(common.Rate/common.TicketLength) * time.Millisecond)
		}
	}
}

func sendOneReq() {
	resp, err := clientGRPC.SayHello(context.Background(), &pb.HelloRequest{Message: common.MessageBody})
	if err != nil {
		log.Println(err)
	}
	log.Println("Resp grpc: ", resp.Message)
}
