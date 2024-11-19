package grpc

import (
	"client/common"
	pb "client/producer/grpc/proto"
	"context"
	"time"
)

func sendReq() {
	for {
		if !isSend {
			time.Sleep(10 * time.Second)
			continue
		}
		clientGRPC.SayHello(context.Background(), &pb.HelloRequest{Message: common.MessageBody})
	}
}

func sendOneReq() {
	clientGRPC.SayHello(context.Background(), &pb.HelloRequest{Message: common.MessageBody})
}
