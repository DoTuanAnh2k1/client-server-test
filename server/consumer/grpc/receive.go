package grpc

import (
	"context"
	"server/common"
	pb "server/consumer/grpc/proto"
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	common.CountRequestStart++
	if req.Message == common.MessageBody {
		common.CountRequestSuccess++
	}
	return &pb.HelloResponse{Message: "Hello " + req.Message}, nil
}
