package main

import (
	"context"
	"fmt"
	pb "grpc-test/proto"
	"io"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.Message}, nil
}

func (s *server) SayHelloServerStream(req *pb.HelloRequest, stream pb.MyService_SayHelloServerStreamServer) error {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		resp := &pb.HelloResponse{Message: fmt.Sprintf("Hello %s %d", req.Message, i)}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) SayHelloClientStream(stream pb.MyService_SayHelloClientStreamServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				response := &pb.HelloResponse{Message: fmt.Sprintf("Hello to all: %v", messages)}
				return stream.SendAndClose(response)
			}
			return err
		}
		messages = append(messages, req.Message)
	}
}

func (s *server) SayHelloBidirectional(stream pb.MyService_SayHelloBidirectionalServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		resp := &pb.HelloResponse{Message: "Hello " + req.Message}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}

// In case cannot gen proto
// export GO_PATH=~/go
// export PATH=$PATH:/$GO_PATH/bin
// protoc --go_out=. --go-grpc_out=. proto/mess.proto

const (
	GPRCTypeUnary        uint8 = 0
	GPRCTypeServerStream uint8 = 1
	GPRCTypeClientStream uint8 = 2
	GPRCTypeBiStream     uint8 = 3
)

func NewClient(gRPCType uint8) {
	url := "localhost:50051"
	clientConnection, err := grpc.NewClient(url, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	clientGRPC := pb.NewMyServiceClient(clientConnection)
	switch gRPCType {
	case GPRCTypeUnary:
		ClientUnary(clientGRPC)
	case GPRCTypeServerStream:
		ClientServerStream(clientGRPC)
	case GPRCTypeClientStream:
		ClientClientStream(clientGRPC)
	case GPRCTypeBiStream:
		ClientBidirectional(clientGRPC)
	default:
		panic("wrong gRPC type")
	}
}

func ClientUnary(client pb.MyServiceClient) {
	req := &pb.HelloRequest{Message: "unary"}
	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message)
}

func ClientServerStream(client pb.MyServiceClient) {
	req := &pb.HelloRequest{Message: "server streaming"}
	stream, err := client.SayHelloServerStream(context.Background(), req)
	if err != nil {
		panic(err)
	}
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(resp.Message)
	}
}

func ClientClientStream(client pb.MyServiceClient) {
	stream, err := client.SayHelloClientStream(context.Background())
	if err != nil {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		req := &pb.HelloRequest{Message: fmt.Sprintf("client message %d", i)}
		if err := stream.Send(req); err != nil {
			panic(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message)
}

func ClientBidirectional(client pb.MyServiceClient) {
	stream, err := client.SayHelloBidirectional(context.Background())
	if err != nil {
		panic(err)
	}
	go func() {
		for i := 0; i < 5; i++ {
			req := &pb.HelloRequest{Message: fmt.Sprintf("client message %d", i)}
			if err := stream.Send(req); err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(resp.Message)
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
	// NewClient(0)
	// NewClient(1)
	// NewClient(2)
	NewClient(3)
}
