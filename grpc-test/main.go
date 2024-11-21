package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-test/proto"
	"net"
	"time"
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.Message}, nil
}

// In case cannot gen proto
// export GO_PATH=~/go
// export PATH=$PATH:/$GO_PATH/bin

func NewClient() {
	url := "http://localhost:50051"
	clientConnection, err := grpc.NewClient(url)
	if err != nil {
		panic(err)
	}

	clientGRPC := pb.NewMyServiceClient(clientConnection)

	mess := "message"

	for {
		time.Sleep(2 * time.Second)
		go func() {
			resp, err := clientGRPC.SayHello(context.Background(), &pb.HelloRequest{Message: mess})
			if err != nil {
				panic(err)
			}
			fmt.Println(resp.Message)
		}()
	}
}

func NewServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	serverGRPC := grpc.NewServer()
	pb.RegisterMyServiceServer(serverGRPC, &server{})

	if err := serverGRPC.Serve(lis); err != nil {
		panic(err)
	}
}

func main() {
	go NewServer()
	NewClient()
}