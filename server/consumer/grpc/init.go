package grpc

import (
	"net"
	pb "server/consumer/grpc/proto"
	"server/utils"

	"google.golang.org/grpc"
)

func Init() {
	serverGRPCPort = utils.GetEnv("ServerGRPCPort", "50051")
	go InitServerGRPC()
}

func InitServerGRPC() error {
	lis, err := net.Listen("tcp", ":"+serverGRPCPort)
	if err != nil {
		panic(err)
	}

	serverGRPC = grpc.NewServer()
	pb.RegisterMyServiceServer(serverGRPC, &server{})

	if err := serverGRPC.Serve(lis); err != nil {
		panic(err)
	}
	return nil
}
