// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: fulcrum.proto

package fulcrum_service

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

const (
	InformanteWrite_ExecuteCommand_FullMethodName = "/fulcrum_service.InformanteWrite/ExecuteCommand"
)

// InformanteWriteClient is the client API for InformanteWrite service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InformanteWriteClient interface {
	ExecuteCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error)
}

type informanteWriteClient struct {
	cc grpc.ClientConnInterface
}

func NewInformanteWriteClient(cc grpc.ClientConnInterface) InformanteWriteClient {
	return &informanteWriteClient{cc}
}

func (c *informanteWriteClient) ExecuteCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error) {
	out := new(CommandReply)
	err := c.cc.Invoke(ctx, InformanteWrite_ExecuteCommand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InformanteWriteServer is the server API for InformanteWrite service.
// All implementations must embed UnimplementedInformanteWriteServer
// for forward compatibility
type InformanteWriteServer interface {
	ExecuteCommand(context.Context, *CommandRequest) (*CommandReply, error)
	mustEmbedUnimplementedInformanteWriteServer()
}

// UnimplementedInformanteWriteServer must be embedded to have forward compatible implementations.
type UnimplementedInformanteWriteServer struct {
}

func (UnimplementedInformanteWriteServer) ExecuteCommand(context.Context, *CommandRequest) (*CommandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteCommand not implemented")
}
func (UnimplementedInformanteWriteServer) mustEmbedUnimplementedInformanteWriteServer() {}

// UnsafeInformanteWriteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InformanteWriteServer will
// result in compilation errors.
type UnsafeInformanteWriteServer interface {
	mustEmbedUnimplementedInformanteWriteServer()
}

func RegisterInformanteWriteServer(s grpc.ServiceRegistrar, srv InformanteWriteServer) {
	s.RegisterService(&InformanteWrite_ServiceDesc, srv)
}

func _InformanteWrite_ExecuteCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformanteWriteServer).ExecuteCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InformanteWrite_ExecuteCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformanteWriteServer).ExecuteCommand(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InformanteWrite_ServiceDesc is the grpc.ServiceDesc for InformanteWrite service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InformanteWrite_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fulcrum_service.InformanteWrite",
	HandlerType: (*InformanteWriteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteCommand",
			Handler:    _InformanteWrite_ExecuteCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fulcrum.proto",
}

const (
	BrokerRequestFulcrum_EnemyInformationBroker_FullMethodName = "/fulcrum_service.BrokerRequestFulcrum/EnemyInformationBroker"
)

// BrokerRequestFulcrumClient is the client API for BrokerRequestFulcrum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrokerRequestFulcrumClient interface {
	EnemyInformationBroker(ctx context.Context, in *EnemyRequestBroker, opts ...grpc.CallOption) (*EnemyReplyBroker, error)
}

type brokerRequestFulcrumClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerRequestFulcrumClient(cc grpc.ClientConnInterface) BrokerRequestFulcrumClient {
	return &brokerRequestFulcrumClient{cc}
}

func (c *brokerRequestFulcrumClient) EnemyInformationBroker(ctx context.Context, in *EnemyRequestBroker, opts ...grpc.CallOption) (*EnemyReplyBroker, error) {
	out := new(EnemyReplyBroker)
	err := c.cc.Invoke(ctx, BrokerRequestFulcrum_EnemyInformationBroker_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerRequestFulcrumServer is the server API for BrokerRequestFulcrum service.
// All implementations must embed UnimplementedBrokerRequestFulcrumServer
// for forward compatibility
type BrokerRequestFulcrumServer interface {
	EnemyInformationBroker(context.Context, *EnemyRequestBroker) (*EnemyReplyBroker, error)
	mustEmbedUnimplementedBrokerRequestFulcrumServer()
}

// UnimplementedBrokerRequestFulcrumServer must be embedded to have forward compatible implementations.
type UnimplementedBrokerRequestFulcrumServer struct {
}

func (UnimplementedBrokerRequestFulcrumServer) EnemyInformationBroker(context.Context, *EnemyRequestBroker) (*EnemyReplyBroker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnemyInformationBroker not implemented")
}
func (UnimplementedBrokerRequestFulcrumServer) mustEmbedUnimplementedBrokerRequestFulcrumServer() {}

// UnsafeBrokerRequestFulcrumServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerRequestFulcrumServer will
// result in compilation errors.
type UnsafeBrokerRequestFulcrumServer interface {
	mustEmbedUnimplementedBrokerRequestFulcrumServer()
}

func RegisterBrokerRequestFulcrumServer(s grpc.ServiceRegistrar, srv BrokerRequestFulcrumServer) {
	s.RegisterService(&BrokerRequestFulcrum_ServiceDesc, srv)
}

func _BrokerRequestFulcrum_EnemyInformationBroker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnemyRequestBroker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerRequestFulcrumServer).EnemyInformationBroker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BrokerRequestFulcrum_EnemyInformationBroker_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerRequestFulcrumServer).EnemyInformationBroker(ctx, req.(*EnemyRequestBroker))
	}
	return interceptor(ctx, in, info, handler)
}

// BrokerRequestFulcrum_ServiceDesc is the grpc.ServiceDesc for BrokerRequestFulcrum service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrokerRequestFulcrum_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fulcrum_service.BrokerRequestFulcrum",
	HandlerType: (*BrokerRequestFulcrumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnemyInformationBroker",
			Handler:    _BrokerRequestFulcrum_EnemyInformationBroker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fulcrum.proto",
}
