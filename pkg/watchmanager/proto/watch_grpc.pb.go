// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: watch.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WatchClient is the client API for Watch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatchClient interface {
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (Watch_SubscribeClient, error)
}

type watchClient struct {
	cc grpc.ClientConnInterface
}

func NewWatchClient(cc grpc.ClientConnInterface) WatchClient {
	return &watchClient{cc}
}

func (c *watchClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (Watch_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Watch_ServiceDesc.Streams[0], "/central.events.v1.watch/subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchSubscribeClient{stream}
	return x, nil
}

type Watch_SubscribeClient interface {
	Send(*Request) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type watchSubscribeClient struct {
	grpc.ClientStream
}

func (x *watchSubscribeClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *watchSubscribeClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WatchServer is the server API for Watch service.
// All implementations must embed UnimplementedWatchServer
// for forward compatibility
type WatchServer interface {
	Subscribe(Watch_SubscribeServer) error
	mustEmbedUnimplementedWatchServer()
}

// UnimplementedWatchServer must be embedded to have forward compatible implementations.
type UnimplementedWatchServer struct {
}

func (UnimplementedWatchServer) Subscribe(Watch_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedWatchServer) mustEmbedUnimplementedWatchServer() {}

// UnsafeWatchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatchServer will
// result in compilation errors.
type UnsafeWatchServer interface {
	mustEmbedUnimplementedWatchServer()
}

func RegisterWatchServer(s grpc.ServiceRegistrar, srv WatchServer) {
	s.RegisterService(&Watch_ServiceDesc, srv)
}

func _Watch_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WatchServer).Subscribe(&watchSubscribeServer{stream})
}

type Watch_SubscribeServer interface {
	Send(*Event) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type watchSubscribeServer struct {
	grpc.ServerStream
}

func (x *watchSubscribeServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *watchSubscribeServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Watch_ServiceDesc is the grpc.ServiceDesc for Watch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Watch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "central.events.v1.watch",
	HandlerType: (*WatchServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "subscribe",
			Handler:       _Watch_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "watch.proto",
}
