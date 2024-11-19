package grpc

import "client/producer/grpc/proto"

var clientGRPC proto.MyServiceClient

var isSend bool

var (
	serverSvc      string
	serverGRPCPort string
)
