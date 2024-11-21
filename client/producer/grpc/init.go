package grpc

import (
	pb "client/producer/grpc/proto"
	"client/utils"
	"log"

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
	log.Println("Init gRPC connection to: ", url)
	clientConnection, err := grpc.NewClient(url, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return err
	}

	clientGRPC = pb.NewMyServiceClient(clientConnection)
	return nil
}
