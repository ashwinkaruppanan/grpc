// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: proto/email.proto

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

// EmailGeneratorClient is the client API for EmailGenerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailGeneratorClient interface {
	UnaryConnection(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Email, error)
	ServerStreamConnection(ctx context.Context, in *NameList, opts ...grpc.CallOption) (EmailGenerator_ServerStreamConnectionClient, error)
	ClientStreamConnection(ctx context.Context, opts ...grpc.CallOption) (EmailGenerator_ClientStreamConnectionClient, error)
	BidirectionalStreamConnection(ctx context.Context, opts ...grpc.CallOption) (EmailGenerator_BidirectionalStreamConnectionClient, error)
}

type emailGeneratorClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailGeneratorClient(cc grpc.ClientConnInterface) EmailGeneratorClient {
	return &emailGeneratorClient{cc}
}

func (c *emailGeneratorClient) UnaryConnection(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Email, error) {
	out := new(Email)
	err := c.cc.Invoke(ctx, "/email.EmailGenerator/UnaryConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailGeneratorClient) ServerStreamConnection(ctx context.Context, in *NameList, opts ...grpc.CallOption) (EmailGenerator_ServerStreamConnectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &EmailGenerator_ServiceDesc.Streams[0], "/email.EmailGenerator/ServerStreamConnection", opts...)
	if err != nil {
		return nil, err
	}
	x := &emailGeneratorServerStreamConnectionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EmailGenerator_ServerStreamConnectionClient interface {
	Recv() (*Email, error)
	grpc.ClientStream
}

type emailGeneratorServerStreamConnectionClient struct {
	grpc.ClientStream
}

func (x *emailGeneratorServerStreamConnectionClient) Recv() (*Email, error) {
	m := new(Email)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *emailGeneratorClient) ClientStreamConnection(ctx context.Context, opts ...grpc.CallOption) (EmailGenerator_ClientStreamConnectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &EmailGenerator_ServiceDesc.Streams[1], "/email.EmailGenerator/ClientStreamConnection", opts...)
	if err != nil {
		return nil, err
	}
	x := &emailGeneratorClientStreamConnectionClient{stream}
	return x, nil
}

type EmailGenerator_ClientStreamConnectionClient interface {
	Send(*Name) error
	CloseAndRecv() (*EmailList, error)
	grpc.ClientStream
}

type emailGeneratorClientStreamConnectionClient struct {
	grpc.ClientStream
}

func (x *emailGeneratorClientStreamConnectionClient) Send(m *Name) error {
	return x.ClientStream.SendMsg(m)
}

func (x *emailGeneratorClientStreamConnectionClient) CloseAndRecv() (*EmailList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmailList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *emailGeneratorClient) BidirectionalStreamConnection(ctx context.Context, opts ...grpc.CallOption) (EmailGenerator_BidirectionalStreamConnectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &EmailGenerator_ServiceDesc.Streams[2], "/email.EmailGenerator/BidirectionalStreamConnection", opts...)
	if err != nil {
		return nil, err
	}
	x := &emailGeneratorBidirectionalStreamConnectionClient{stream}
	return x, nil
}

type EmailGenerator_BidirectionalStreamConnectionClient interface {
	Send(*Name) error
	Recv() (*Email, error)
	grpc.ClientStream
}

type emailGeneratorBidirectionalStreamConnectionClient struct {
	grpc.ClientStream
}

func (x *emailGeneratorBidirectionalStreamConnectionClient) Send(m *Name) error {
	return x.ClientStream.SendMsg(m)
}

func (x *emailGeneratorBidirectionalStreamConnectionClient) Recv() (*Email, error) {
	m := new(Email)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EmailGeneratorServer is the server API for EmailGenerator service.
// All implementations must embed UnimplementedEmailGeneratorServer
// for forward compatibility
type EmailGeneratorServer interface {
	UnaryConnection(context.Context, *Name) (*Email, error)
	ServerStreamConnection(*NameList, EmailGenerator_ServerStreamConnectionServer) error
	ClientStreamConnection(EmailGenerator_ClientStreamConnectionServer) error
	BidirectionalStreamConnection(EmailGenerator_BidirectionalStreamConnectionServer) error
	mustEmbedUnimplementedEmailGeneratorServer()
}

// UnimplementedEmailGeneratorServer must be embedded to have forward compatible implementations.
type UnimplementedEmailGeneratorServer struct {
}

func (UnimplementedEmailGeneratorServer) UnaryConnection(context.Context, *Name) (*Email, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryConnection not implemented")
}
func (UnimplementedEmailGeneratorServer) ServerStreamConnection(*NameList, EmailGenerator_ServerStreamConnectionServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamConnection not implemented")
}
func (UnimplementedEmailGeneratorServer) ClientStreamConnection(EmailGenerator_ClientStreamConnectionServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamConnection not implemented")
}
func (UnimplementedEmailGeneratorServer) BidirectionalStreamConnection(EmailGenerator_BidirectionalStreamConnectionServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreamConnection not implemented")
}
func (UnimplementedEmailGeneratorServer) mustEmbedUnimplementedEmailGeneratorServer() {}

// UnsafeEmailGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailGeneratorServer will
// result in compilation errors.
type UnsafeEmailGeneratorServer interface {
	mustEmbedUnimplementedEmailGeneratorServer()
}

func RegisterEmailGeneratorServer(s grpc.ServiceRegistrar, srv EmailGeneratorServer) {
	s.RegisterService(&EmailGenerator_ServiceDesc, srv)
}

func _EmailGenerator_UnaryConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailGeneratorServer).UnaryConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email.EmailGenerator/UnaryConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailGeneratorServer).UnaryConnection(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailGenerator_ServerStreamConnection_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NameList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EmailGeneratorServer).ServerStreamConnection(m, &emailGeneratorServerStreamConnectionServer{stream})
}

type EmailGenerator_ServerStreamConnectionServer interface {
	Send(*Email) error
	grpc.ServerStream
}

type emailGeneratorServerStreamConnectionServer struct {
	grpc.ServerStream
}

func (x *emailGeneratorServerStreamConnectionServer) Send(m *Email) error {
	return x.ServerStream.SendMsg(m)
}

func _EmailGenerator_ClientStreamConnection_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmailGeneratorServer).ClientStreamConnection(&emailGeneratorClientStreamConnectionServer{stream})
}

type EmailGenerator_ClientStreamConnectionServer interface {
	SendAndClose(*EmailList) error
	Recv() (*Name, error)
	grpc.ServerStream
}

type emailGeneratorClientStreamConnectionServer struct {
	grpc.ServerStream
}

func (x *emailGeneratorClientStreamConnectionServer) SendAndClose(m *EmailList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *emailGeneratorClientStreamConnectionServer) Recv() (*Name, error) {
	m := new(Name)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EmailGenerator_BidirectionalStreamConnection_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmailGeneratorServer).BidirectionalStreamConnection(&emailGeneratorBidirectionalStreamConnectionServer{stream})
}

type EmailGenerator_BidirectionalStreamConnectionServer interface {
	Send(*Email) error
	Recv() (*Name, error)
	grpc.ServerStream
}

type emailGeneratorBidirectionalStreamConnectionServer struct {
	grpc.ServerStream
}

func (x *emailGeneratorBidirectionalStreamConnectionServer) Send(m *Email) error {
	return x.ServerStream.SendMsg(m)
}

func (x *emailGeneratorBidirectionalStreamConnectionServer) Recv() (*Name, error) {
	m := new(Name)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EmailGenerator_ServiceDesc is the grpc.ServiceDesc for EmailGenerator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailGenerator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "email.EmailGenerator",
	HandlerType: (*EmailGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryConnection",
			Handler:    _EmailGenerator_UnaryConnection_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamConnection",
			Handler:       _EmailGenerator_ServerStreamConnection_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamConnection",
			Handler:       _EmailGenerator_ClientStreamConnection_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStreamConnection",
			Handler:       _EmailGenerator_BidirectionalStreamConnection_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/email.proto",
}
