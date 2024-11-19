package grpc

import (
	pb "client/producer/grpc/proto"
	"client/utils"

	"google.golang.org/grpc"
)

// In case cannot gen proto
// export GO_PATH=~/go
// export PATH=$PATH:/$GO_PATH/bin

func Init() error {
	serverSvc = utils.GetEnv("ServerSvc", "localhost")
	serverGRPCPort = utils.GetEnv("ServerGRPCPort", "50051")
	err := InitClient()
	if err != nil {
		return err
	}
	go sendReq()
	return nil
}

func InitClient() error {
	url := serverSvc + ":" + serverGRPCPort
	clientConnection, err := grpc.NewClient(url)
	if err != nil {
		return nil
	}

	clientGRPC = pb.NewMyServiceClient(clientConnection)
	return nil
}
