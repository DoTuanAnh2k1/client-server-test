// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: proto/mess.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MyService_SayHello_FullMethodName              = "/proto.MyService/SayHello"
	MyService_SayHelloServerStream_FullMethodName  = "/proto.MyService/SayHelloServerStream"
	MyService_SayHelloClientStream_FullMethodName  = "/proto.MyService/SayHelloClientStream"
	MyService_SayHelloBidirectional_FullMethodName = "/proto.MyService/SayHelloBidirectional"
)

// MyServiceClient is the client API for MyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	SayHelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HelloResponse], error)
	SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[HelloRequest, HelloResponse], error)
	SayHelloBidirectional(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[HelloRequest, HelloResponse], error)
}

type myServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyServiceClient(cc grpc.ClientConnInterface) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, MyService_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myServiceClient) SayHelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HelloResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MyService_ServiceDesc.Streams[0], MyService_SayHelloServerStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HelloRequest, HelloResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MyService_SayHelloServerStreamClient = grpc.ServerStreamingClient[HelloResponse]

func (c *myServiceClient) SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[HelloRequest, HelloResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MyService_ServiceDesc.Streams[1], MyService_SayHelloClientStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HelloRequest, HelloResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MyService_SayHelloClientStreamClient = grpc.ClientStreamingClient[HelloRequest, HelloResponse]

func (c *myServiceClient) SayHelloBidirectional(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[HelloRequest, HelloResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MyService_ServiceDesc.Streams[2], MyService_SayHelloBidirectional_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HelloRequest, HelloResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MyService_SayHelloBidirectionalClient = grpc.BidiStreamingClient[HelloRequest, HelloResponse]

// MyServiceServer is the server API for MyService service.
// All implementations must embed UnimplementedMyServiceServer
// for forward compatibility.
type MyServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	SayHelloServerStream(*HelloRequest, grpc.ServerStreamingServer[HelloResponse]) error
	SayHelloClientStream(grpc.ClientStreamingServer[HelloRequest, HelloResponse]) error
	SayHelloBidirectional(grpc.BidiStreamingServer[HelloRequest, HelloResponse]) error
	mustEmbedUnimplementedMyServiceServer()
}

// UnimplementedMyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMyServiceServer struct{}

func (UnimplementedMyServiceServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedMyServiceServer) SayHelloServerStream(*HelloRequest, grpc.ServerStreamingServer[HelloResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStream not implemented")
}
func (UnimplementedMyServiceServer) SayHelloClientStream(grpc.ClientStreamingServer[HelloRequest, HelloResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloClientStream not implemented")
}
func (UnimplementedMyServiceServer) SayHelloBidirectional(grpc.BidiStreamingServer[HelloRequest, HelloResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloBidirectional not implemented")
}
func (UnimplementedMyServiceServer) mustEmbedUnimplementedMyServiceServer() {}
func (UnimplementedMyServiceServer) testEmbeddedByValue()                   {}

// UnsafeMyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyServiceServer will
// result in compilation errors.
type UnsafeMyServiceServer interface {
	mustEmbedUnimplementedMyServiceServer()
}

func RegisterMyServiceServer(s grpc.ServiceRegistrar, srv MyServiceServer) {
	// If the following call pancis, it indicates UnimplementedMyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MyService_ServiceDesc, srv)
}

func _MyService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MyService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyService_SayHelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MyServiceServer).SayHelloServerStream(m, &grpc.GenericServerStream[HelloRequest, HelloResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MyService_SayHelloServerStreamServer = grpc.ServerStreamingServer[HelloResponse]

func _MyService_SayHelloClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MyServiceServer).SayHelloClientStream(&grpc.GenericServerStream[HelloRequest, HelloResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MyService_SayHelloClientStreamServer = grpc.ClientStreamingServer[HelloRequest, HelloResponse]

func _MyService_SayHelloBidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MyServiceServer).SayHelloBidirectional(&grpc.GenericServerStream[HelloRequest, HelloResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MyService_SayHelloBidirectionalServer = grpc.BidiStreamingServer[HelloRequest, HelloResponse]

// MyService_ServiceDesc is the grpc.ServiceDesc for MyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _MyService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloServerStream",
			Handler:       _MyService_SayHelloServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloClientStream",
			Handler:       _MyService_SayHelloClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloBidirectional",
			Handler:       _MyService_SayHelloBidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/mess.proto",
}
