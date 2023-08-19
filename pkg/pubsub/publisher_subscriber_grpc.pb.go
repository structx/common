// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: protos/publisher_subscriber.proto

package pubsub_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PubSubServiceClient is the client API for PubSubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubSubServiceClient interface {
	Publish(ctx context.Context, opts ...grpc.CallOption) (PubSubService_PublishClient, error)
	Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (PubSubService_SubscribeClient, error)
}

type pubSubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPubSubServiceClient(cc grpc.ClientConnInterface) PubSubServiceClient {
	return &pubSubServiceClient{cc}
}

func (c *pubSubServiceClient) Publish(ctx context.Context, opts ...grpc.CallOption) (PubSubService_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &PubSubService_ServiceDesc.Streams[0], "/pubsub.PubSubService/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubSubServicePublishClient{stream}
	return x, nil
}

type PubSubService_PublishClient interface {
	Send(*Message) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type pubSubServicePublishClient struct {
	grpc.ClientStream
}

func (x *pubSubServicePublishClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pubSubServicePublishClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pubSubServiceClient) Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (PubSubService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &PubSubService_ServiceDesc.Streams[1], "/pubsub.PubSubService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubSubServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubSubService_SubscribeClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type pubSubServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubSubServiceSubscribeClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PubSubServiceServer is the server API for PubSubService service.
// All implementations must embed UnimplementedPubSubServiceServer
// for forward compatibility
type PubSubServiceServer interface {
	Publish(PubSubService_PublishServer) error
	Subscribe(*Subscription, PubSubService_SubscribeServer) error
	mustEmbedUnimplementedPubSubServiceServer()
}

// UnimplementedPubSubServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPubSubServiceServer struct {
}

func (UnimplementedPubSubServiceServer) Publish(PubSubService_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedPubSubServiceServer) Subscribe(*Subscription, PubSubService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedPubSubServiceServer) mustEmbedUnimplementedPubSubServiceServer() {}

// UnsafePubSubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubSubServiceServer will
// result in compilation errors.
type UnsafePubSubServiceServer interface {
	mustEmbedUnimplementedPubSubServiceServer()
}

func RegisterPubSubServiceServer(s grpc.ServiceRegistrar, srv PubSubServiceServer) {
	s.RegisterService(&PubSubService_ServiceDesc, srv)
}

func _PubSubService_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PubSubServiceServer).Publish(&pubSubServicePublishServer{stream})
}

type PubSubService_PublishServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pubSubServicePublishServer struct {
	grpc.ServerStream
}

func (x *pubSubServicePublishServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pubSubServicePublishServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PubSubService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Subscription)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubSubServiceServer).Subscribe(m, &pubSubServiceSubscribeServer{stream})
}

type PubSubService_SubscribeServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type pubSubServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubSubServiceSubscribeServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

// PubSubService_ServiceDesc is the grpc.ServiceDesc for PubSubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubSubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pubsub.PubSubService",
	HandlerType: (*PubSubServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Publish",
			Handler:       _PubSubService_Publish_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _PubSubService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/publisher_subscriber.proto",
}
