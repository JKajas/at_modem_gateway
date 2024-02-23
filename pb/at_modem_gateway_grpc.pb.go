// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: at_modem_gateway.proto

package sms_gateway

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

// SmsGatewayClient is the client API for SmsGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsGatewayClient interface {
	SendSms(ctx context.Context, in *GtwRequest, opts ...grpc.CallOption) (*GtwResponse, error)
}

type smsGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsGatewayClient(cc grpc.ClientConnInterface) SmsGatewayClient {
	return &smsGatewayClient{cc}
}

func (c *smsGatewayClient) SendSms(ctx context.Context, in *GtwRequest, opts ...grpc.CallOption) (*GtwResponse, error) {
	out := new(GtwResponse)
	err := c.cc.Invoke(ctx, "/SmsGateway/SendSms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsGatewayServer is the server API for SmsGateway service.
// All implementations must embed UnimplementedSmsGatewayServer
// for forward compatibility
type SmsGatewayServer interface {
	SendSms(context.Context, *GtwRequest) (*GtwResponse, error)
	mustEmbedUnimplementedSmsGatewayServer()
}

// UnimplementedSmsGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedSmsGatewayServer struct {
}

func (UnimplementedSmsGatewayServer) SendSms(context.Context, *GtwRequest) (*GtwResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSms not implemented")
}
func (UnimplementedSmsGatewayServer) mustEmbedUnimplementedSmsGatewayServer() {}

// UnsafeSmsGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsGatewayServer will
// result in compilation errors.
type UnsafeSmsGatewayServer interface {
	mustEmbedUnimplementedSmsGatewayServer()
}

func RegisterSmsGatewayServer(s grpc.ServiceRegistrar, srv SmsGatewayServer) {
	s.RegisterService(&SmsGateway_ServiceDesc, srv)
}

func _SmsGateway_SendSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GtwRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsGatewayServer).SendSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmsGateway/SendSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsGatewayServer).SendSms(ctx, req.(*GtwRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmsGateway_ServiceDesc is the grpc.ServiceDesc for SmsGateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmsGateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SmsGateway",
	HandlerType: (*SmsGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSms",
			Handler:    _SmsGateway_SendSms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "at_modem_gateway.proto",
}
