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
		_, err := clientGRPC.SayHello(context.Background(), &pb.HelloRequest{Message: common.MessageBody})
		if err != nil {
			log.Fatalf("Error during call: %v", err)
		}
	}
}
